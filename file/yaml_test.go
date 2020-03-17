package file

import (
	"testing"
)

type ServerConfig struct {
	Addr string
	Timeout string
}

func TestLoad(t *testing.T) {
	tests := []struct{
		key string
		val interface{}
	} {
		{"Server", ServerConfig{ "0.0.0.0:8080", "1s"}},
		{"Client", ServerConfig{ "0.0.0.0:8000", "1s"}},
		{"Test", "test"},
		{"TestNum", 1},
		{"alibaba", nil},
		{"test", nil},
	}

	y, err := Load("test.yaml")
	if err != nil {
		t.Error(err)
	}

	for _, v := range tests {
		if v.key == "Server" || v.key == "Client" {
			var cfg ServerConfig
			y.Unmarshal(v.key, &cfg)
			if cfg != v.val {
				t.Errorf("%s was %s, but it %s \n", v.key, cfg, v.val)
			}
			continue
		}

		var val interface{}
		y.Unmarshal(v.key, &val)
		if val != v.val {
			t.Errorf("%s was %s, but it %s \n", v.key, val, v.val)
		}
	}
}
