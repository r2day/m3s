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
