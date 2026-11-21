package actor

import (
	"encoding/json"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// transport-api 按 "_.meta.merchant_id"、"account_id"、"driver.status"、"wx_openid" 查询和建索引。
func TestBSONFieldNames(t *testing.T) {
	m := &Model{AccountID: "a1", WxOpenID: "o1", Roles: []string{"driver"}, Driver: Driver{Status: "pending"}}
	m.Meta.MerchantID = "m1"
	raw, err := bson.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var doc bson.M
	if err := bson.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	meta, _ := doc["_"].(bson.M)["meta"].(bson.M)
	if meta["merchant_id"] != "m1" || doc["account_id"] != "a1" || doc["wx_openid"] != "o1" {
		t.Fatalf("%v", doc)
	}
	if d, _ := doc["driver"].(bson.M); d["status"] != "pending" {
		t.Fatalf("driver=%v", doc["driver"])
	}
}

// openid 只存库，不能被序列化给小程序。
func TestJSONHidesOpenID(t *testing.T) {
	raw, err := json.Marshal(&Model{WxOpenID: "secret-openid"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "secret-openid") || strings.Contains(string(raw), "wx_openid") {
		t.Fatal(string(raw))
	}
}
