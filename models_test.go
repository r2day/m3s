package m3s_test

import (
	"encoding/json"
	"testing"

	home "github.com/r2day/m3s/client/home" // 该目录的包名是 notification
	"github.com/r2day/m3s/client/notification"
	"github.com/r2day/m3s/client/theme"
	"github.com/r2day/m3s/device/printer"
	"github.com/r2day/m3s/device/ptpl"
	"github.com/r2day/m3s/f5s/file"
	"github.com/r2day/m3s/finance/keys"
	"github.com/r2day/m3s/message/subscribe"
	"github.com/r2day/m3s/stat/monthly"
	"github.com/r2day/m3s/transport/address"
	"github.com/r2day/m3s/transport/fleet"
	"github.com/r2day/m3s/transport/order"
	"github.com/r2day/m3s/transport/platform"
	"github.com/r2day/m3s/transport/route"
	"github.com/r2day/m3s/transport/vehicle"
	"go.mongodb.org/mongo-driver/bson"
)

type named interface {
	ResourceName() string
	CollectionName() string
}

// 集合名是各服务共用的契约，改名会让线上数据"消失"，这里逐个固定下来。
func TestResourceAndCollectionNames(t *testing.T) {
	cases := []struct {
		name               string
		m                  named
		resource, collName string
	}{
		{"client/home", &home.Model{}, "home", "client_home_config"},
		{"client/notification", &notification.Model{}, "notification", "client_notification_config"},
		{"device/printer", &printer.Model{}, "printer", "device_printer_config"},
		{"device/ptpl", &ptpl.Model{}, "printer", "device_printer_template"},
		{"f5s/file", &file.Model{}, "file", "f5s_file_flow"},
		{"finance/keys", &keys.Model{}, "keys", "finance_keys_config"},
		{"message/subscribe", &subscribe.Model{}, "subscribe", "message_subscribe_flow"},
		{"stat/monthly", &monthly.Model{}, "monthly", "stat_monthly_data"},
		{"transport/address", &address.Model{}, "address", "transport_address_config"},
		{"transport/fleet", &fleet.Model{}, "fleet", "transport_fleet_config"},
		{"transport/order", &order.Model{}, "takeout", "transport_takeout_order"},
		{"transport/platform", &platform.Model{}, "platform", "transport_platform_config"},
		{"transport/route", &route.Model{}, "route", "transport_route_config"},
		{"transport/vehicle", &vehicle.Model{}, "vehicle", "transport_vehicle_config"},
	}
	seen := map[string]string{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.m.ResourceName(); got != tc.resource {
				t.Fatalf("ResourceName=%q want %q", got, tc.resource)
			}
			if got := tc.m.CollectionName(); got != tc.collName {
				t.Fatalf("CollectionName=%q want %q", got, tc.collName)
			}
		})
		if prev, dup := seen[tc.collName]; dup {
			t.Fatalf("集合名重复: %s 与 %s", tc.name, prev)
		}
		seen[tc.collName] = tc.name
	}
	if got := (&theme.Model{}).CollectionName(); got != "client_theme_template" {
		t.Fatalf("theme CollectionName=%q", got)
	}
}

// 枚举值会写进数据库并被 fileserver 等服务按数字使用，固定下来避免插入新值导致错位。
func TestEnumValues(t *testing.T) {
	cases := []struct {
		name      string
		got, want int
	}{
		{"file.Default", int(file.Default), 0},
		{"file.Icon", int(file.Icon), 1},
		{"file.Avatar", int(file.Avatar), 4},
		{"file.Logo", int(file.Logo), 7},
		{"file.Motion", int(file.Motion), 8},
		{"file.QRCode", int(file.QRCode), 9},
		{"notification.Sent", int(notification.Sent), 0},
		{"notification.Deleted", int(notification.Deleted), 3},
		{"notification.PrivateMessage", int(notification.PrivateMessage), 2},
		{"home.WxApp", int(home.WxApp), 0},
		{"home.DouYin", int(home.DouYin), 2},
	}
	for _, tc := range cases {
		if tc.got != tc.want {
			t.Fatalf("%s=%d want %d", tc.name, tc.got, tc.want)
		}
	}
}

// 记录现状：内嵌的 model.Model 带 bson:"_"，所以 meta 存在 "_.meta" 下，而不是顶层 "meta"。
// 这正是 GetByStoreID 查询 "meta.merchant_id" 永远查不到数据的原因（见集成测试）。
func TestBSONLayoutPutsMetaUnderUnderscore(t *testing.T) {
	p := &printer.Model{Name: "前台"}
	p.Meta.MerchantID = "m1"
	f := &file.Model{Name: "logo", ContentMD5: "abc"}
	f.Meta.MerchantID = "m1"
	k := &keys.Model{Name: "wx"}
	k.Meta.MerchantID = "m1"
	cases := []struct {
		name string
		m    any
	}{{"printer", p}, {"file", f}, {"keys", k}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			raw, err := bson.Marshal(tc.m)
			if err != nil {
				t.Fatal(err)
			}
			var doc bson.M
			if err := bson.Unmarshal(raw, &doc); err != nil {
				t.Fatal(err)
			}
			if _, ok := doc["meta"]; ok {
				t.Fatal("meta 不应出现在顶层")
			}
			under, ok := doc["_"].(bson.M)
			if !ok {
				t.Fatalf("缺少 _ 字段: %v", doc)
			}
			meta, _ := under["meta"].(bson.M)
			if meta["merchant_id"] != "m1" {
				t.Fatalf("_.meta=%v", under)
			}
			if _, ok := doc["_id"]; ok {
				t.Fatal("零值 ID 应被 omitempty 省略")
			}
		})
	}
}

// JSON 也是 "_": {"meta": ...}；Context 内部字段全是 json:"-"，不会把数据库句柄序列化出去。
func TestJSONLayout(t *testing.T) {
	m := &file.Model{Name: "logo"}
	m.Meta.MerchantID = "m1"
	raw, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var doc map[string]any
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	under, _ := doc["_"].(map[string]any)
	meta, _ := under["meta"].(map[string]any)
	ctx, _ := under["context"].(map[string]any)
	if meta["merchant_id"] != "m1" || len(ctx) != 0 || doc["name"] != "logo" {
		t.Fatalf("json=%s", raw)
	}
}
