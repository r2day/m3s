package freight

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// transport-api 的条件更新直接写这些路径，字段名是契约。
func TestBSONFieldNames(t *testing.T) {
	m := &Model{
		FreightNo:        "FRT1",
		Status:           "published",
		ShipperAccountID: "s1",
		DriverAccountID:  "d1",
		Origin:           Place{City: "上海"},
		Applications:     []Application{{ID: "x", AccountID: "d1", Status: "pending"}},
		Pay:              Pay{Status: "unpaid", Ledger: "none", ShipperConfirmedAt: 1},
		CreatedAt:        1,
	}
	m.Meta.MerchantID = "m1"
	raw, err := bson.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var doc bson.M
	if err := bson.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"freight_no", "status", "shipper_account_id", "driver_account_id", "origin", "applications", "pay", "created_at"} {
		if _, ok := doc[k]; !ok {
			t.Fatalf("缺少 %s: %v", k, doc)
		}
	}
	if _, ok := doc["idempotency_key"]; ok {
		t.Fatal("空 idempotency_key 不应落库（partial 唯一索引）")
	}
	if doc["origin"].(bson.M)["city"] != "上海" {
		t.Fatal("origin.city")
	}
	app := doc["applications"].(bson.A)[0].(bson.M)
	if app["id"] != "x" || app["account_id"] != "d1" || app["status"] != "pending" {
		t.Fatalf("app=%v", app)
	}
	pay := doc["pay"].(bson.M)
	if pay["ledger"] != "none" || pay["status"] != "unpaid" || pay["shipper_confirmed_at"] != int64(1) {
		t.Fatalf("pay=%v", pay)
	}
	if _, ok := pay["driver_confirmed_at"]; ok {
		t.Fatal("未确认的一侧不应写 0（ConfirmFilter 认 $exists:false 或 0，两种都行，这里保持缺失）")
	}
}
