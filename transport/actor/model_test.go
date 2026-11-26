package actor

import "testing"

func TestCollectionName(t *testing.T) {
	m := &Model{}
	got := m.CollectionName()
	if got != "transport_actor_config" {
		t.Fatalf("CollectionName() = %q", got)
	}
	if m.ResourceName() != "actor" {
		t.Fatalf("ResourceName() = %q", m.ResourceName())
	}
}
