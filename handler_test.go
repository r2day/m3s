package m3s

import "testing"

// mongo.Connect 只解析 URI、不建立连接，所以这里不会访问任何数据库。
func TestNewSelectsDatabase(t *testing.T) {
	cases := []struct {
		name, env, want string
	}{
		{"默认库名", "", "m3s"},
		{"M3S_DB 覆盖", "m3s_custom", "m3s_custom"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(dbNameKey, tc.env)
			prev := MDB
			t.Cleanup(func() { MDB = prev })
			New("mongodb://127.0.0.1:1")
			if MDB == nil || MDB.Name() != tc.want {
				t.Fatalf("MDB=%v", MDB)
			}
		})
	}
}

func TestNewPanicsOnBadURI(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("非法 URI 应 panic")
		}
	}()
	New("not-a-mongo-uri")
}
