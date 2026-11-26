package m3s_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	home "github.com/r2day/m3s/client/home" // 该目录的包名是 notification
	"github.com/r2day/m3s/client/notification"
	"github.com/r2day/m3s/device/printer"
	"github.com/r2day/m3s/device/ptpl"
	"github.com/r2day/m3s/f5s/file"
	"github.com/r2day/m3s/finance/keys"
	"github.com/r2day/m3s/transport/order"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// itDB 连接本地 MongoDB 的临时库，测试结束后删除。
// 仅在 R2DAY_IT=1 时运行；地址可用 R2DAY_IT_MONGO 覆盖，默认 mongodb://127.0.0.1:27017。
func itDB(t *testing.T) (context.Context, *mongo.Database) {
	t.Helper()
	if os.Getenv("R2DAY_IT") != "1" {
		t.Skip("集成测试：设置 R2DAY_IT=1 并启动本地 MongoDB 后运行")
	}
	uri := os.Getenv("R2DAY_IT_MONGO")
	if uri == "" {
		uri = "mongodb://127.0.0.1:27017"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatalf("无法连接 MongoDB: %v", err)
	}
	db := client.Database("m3s_it")
	t.Cleanup(func() {
		_ = db.Drop(context.Background())
		_ = client.Disconnect(context.Background())
	})
	return ctx, db
}

// GetByStoreID 查询 {"meta.merchant_id": ObjectID}，但模型实际把商户号以字符串存在 "_.meta.merchant_id"。
// 结果：按模型正常写入的数据永远查不到；只有手工构造的顶层 meta + ObjectID 文档才能命中。
func TestIT_GetByStoreIDFieldMismatch(t *testing.T) {
	ctx, db := itDB(t)
	storeID := primitive.NewObjectID()

	cases := []struct {
		name  string
		coll  string
		model any
		query func() (int, error)
	}{
		{"client/home", (&home.Model{}).CollectionName(), &home.Model{}, func() (int, error) {
			m := &home.Model{}
			m.Init(ctx, db, m.CollectionName())
			r, err := m.GetByStoreID(storeID.Hex())
			return len(r), err
		}},
		{"client/notification", (&notification.Model{}).CollectionName(), &notification.Model{}, func() (int, error) {
			m := &notification.Model{}
			m.Init(ctx, db, m.CollectionName())
			r, err := m.GetByStoreID(storeID.Hex())
			return len(r), err
		}},
		{"device/printer", (&printer.Model{}).CollectionName(), &printer.Model{}, func() (int, error) {
			m := &printer.Model{}
			m.Init(ctx, db, m.CollectionName())
			r, err := m.GetByStoreID(storeID.Hex())
			return len(r), err
		}},
		{"device/ptpl", (&ptpl.Model{}).CollectionName(), &ptpl.Model{}, func() (int, error) {
			m := &ptpl.Model{}
			m.Init(ctx, db, m.CollectionName())
			r, err := m.GetByStoreID(storeID.Hex())
			return len(r), err
		}},
		{"finance/keys", (&keys.Model{}).CollectionName(), &keys.Model{}, func() (int, error) {
			m := &keys.Model{}
			m.Init(ctx, db, m.CollectionName())
			r, err := m.GetByStoreID(storeID.Hex())
			return len(r), err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			coll := db.Collection(tc.coll)
			// 1) 模拟业务代码的写法：商户号是字符串，序列化后落在 "_.meta.merchant_id"
			raw, err := bson.Marshal(tc.model)
			if err != nil {
				t.Fatal(err)
			}
			var doc bson.M
			if err := bson.Unmarshal(raw, &doc); err != nil {
				t.Fatal(err)
			}
			doc["_"] = bson.M{"meta": bson.M{"merchant_id": storeID.Hex()}}
			if _, err := coll.InsertOne(ctx, doc); err != nil {
				t.Fatal(err)
			}
			n, err := tc.query()
			if err != nil || n != 0 {
				t.Fatalf("现状应为查不到：n=%d err=%v", n, err)
			}

			// 2) 只有顶层 meta.merchant_id 且为 ObjectID 的文档才会被返回
			if _, err := coll.InsertOne(ctx, bson.M{"meta": bson.M{"merchant_id": storeID}}); err != nil {
				t.Fatal(err)
			}
			n, err = tc.query()
			if err != nil || n != 1 {
				t.Fatalf("n=%d err=%v", n, err)
			}

			// 3) 非法 hex 会被当成全零 ObjectID 查询，不报错
			m := &printer.Model{}
			m.Init(ctx, db, tc.coll)
			r, err := m.GetByStoreID("not-hex")
			if err != nil || len(r) != 0 {
				t.Fatalf("not-hex: %d %v", len(r), err)
			}
		})
	}
}

func TestIT_OrderLookups(t *testing.T) {
	ctx, db := itDB(t)
	seed := &order.Model{OriginID: "origin-1", OrderCode: "code-1"}
	if _, err := db.Collection(seed.CollectionName()).InsertOne(ctx, seed); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name    string
		fn      func(*order.Model) error
		wantErr error
	}{
		{"按 origin_id 命中", func(m *order.Model) error { return m.GetByOriginID("origin-1") }, nil},
		{"按 order_code 命中", func(m *order.Model) error { return m.GetByOrderCode("code-1") }, nil},
		{"origin_id 不存在", func(m *order.Model) error { return m.GetByOriginID("nope") }, mongo.ErrNoDocuments},
		{"order_code 不存在", func(m *order.Model) error { return m.GetByOrderCode("nope") }, mongo.ErrNoDocuments},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &order.Model{}
			m.Init(ctx, db, m.CollectionName())
			err := tc.fn(m)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err=%v want %v", err, tc.wantErr)
			}
			if tc.wantErr == nil && (m.OriginID != "origin-1" || m.OrderCode != "code-1") {
				t.Fatalf("decoded=%+v", m)
			}
		})
	}
}

// IsExistContent 用 Find 拿到游标后没有调用 Next 就 Decode，所以无论文档是否存在都会返回错误。
func TestIT_IsExistContentAlwaysErrors(t *testing.T) {
	ctx, db := itDB(t)
	seed := &file.Model{Name: "logo", ContentMD5: "md5-exists"}
	if _, err := db.Collection(seed.CollectionName()).InsertOne(ctx, seed); err != nil {
		t.Fatal(err)
	}
	cases := []struct{ name, md5 string }{
		{"文档存在（现状：仍然报错）", "md5-exists"},
		{"文档不存在", "md5-missing"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &file.Model{}
			m.Init(ctx, db, m.CollectionName())
			got, err := m.IsExistContent(tc.md5)
			if err == nil || got != nil {
				t.Fatalf("现状应报错: got=%+v err=%v", got, err)
			}
		})
	}
}
