package theme

import "testing"

func TestCollectionName(t *testing.T) {
	m := &Model{}
	got := m.CollectionName()
	if got != "client_theme_template" {
		t.Errorf("CollectionName() = %q, want %q", got, "client_theme_template")
	}
	if got == "client_theme_config" {
		t.Errorf("CollectionName() = %q, must not be client_theme_config", got)
	}
}
