package order

import (
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

// 空参数直接返回 ErrNoDocuments，不会访问数据库（这里未初始化 Handler）。
func TestLookupsRejectEmptyKey(t *testing.T) {
	cases := []struct {
		name string
		fn   func(*Model) error
	}{
		{"GetByOriginID", func(m *Model) error { return m.GetByOriginID("") }},
		{"GetByOrderCode", func(m *Model) error { return m.GetByOrderCode("") }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.fn(&Model{}); !errors.Is(err, mongo.ErrNoDocuments) {
				t.Fatalf("err=%v", err)
			}
		})
	}
}
