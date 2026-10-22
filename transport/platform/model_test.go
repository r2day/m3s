package platform

import "testing"

func TestBound(t *testing.T) {
	cases := []struct {
		name string
		m    *Model
		want bool
	}{
		{"nil", nil, false},
		{"空", &Model{}, false},
		{"只有空白", &Model{AppID: " ", AppKey: "\t", Token: " \n"}, false},
		{"只有 shop_id 不算绑定", &Model{ShopID: "s1", Provider: "meituan"}, false},
		{"app_id", &Model{AppID: "a"}, true},
		{"app_key", &Model{AppKey: "fake-key"}, true},
		{"token", &Model{Token: "fake-token"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.m.Bound(); got != tc.want {
				t.Fatalf("got %v", got)
			}
		})
	}
}
