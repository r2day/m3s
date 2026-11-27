package freight

import "testing"

func TestCollectionName(t *testing.T) {
	m := &Model{}
	got := m.CollectionName()
	if got != "transport_freight_order" {
		t.Fatalf("CollectionName() = %q", got)
	}
	if got == "transport_takeout_order" {
		t.Fatal("freight must not use the UU collection")
	}
	if m.ResourceName() != "freight" {
		t.Fatalf("ResourceName() = %q", m.ResourceName())
	}
}

func TestPlaceRegionFields(t *testing.T) {
	p := Place{
		Name: "林", City: "上海", Address: "世纪大道",
		Province: "上海市", ProvinceCode: "31", CityCode: "3101",
		District: "浦东新区", DistrictCode: "310115",
		Street: "陆家嘴街道", StreetCode: "310115001",
		Door: "1号", Lat: "31.2", Lng: "121.5",
	}
	if p.ProvinceCode != "31" || p.DistrictCode != "310115" || p.StreetCode == "" {
		t.Fatalf("%+v", p)
	}
}
