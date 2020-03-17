package file

import (
	"gopkg.in/yaml.v2"
	"sort"
	"sync/atomic"
)

type YAML struct {
	v atomic.Value
	data map[string]*interface{}
}

// Load: load file to map
func Load(path string) (y *YAML, err error) {
	y = new(YAML)
	buf, err := ReadFile(path)
	if err != nil {
		return
	}
	y.v.Store(buf)
	y.data = make(map[string]*interface{})
	yaml.Unmarshal(buf, &y.data)
	return
}

func (y *YAML) Get(key string) *interface{} {
	return y.data[key]
}

func (y *YAML) Keys() []string {
	var keys []string
	for k, _ := range y.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (y *YAML) Unmarshal(key string, dst interface{}) (err error) {
	v := y.Get(key)
	buf, err := yaml.Marshal(v)
	if err != nil {
		return
	}
	return yaml.Unmarshal(buf, dst)
}