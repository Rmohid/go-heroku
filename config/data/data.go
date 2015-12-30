// Key value store for global configuration data

package data

import (
	"sort"
	"sync"
)

var (
	data map[string]string
	mu   sync.Mutex
)

func init() {
	data = make(map[string]string)
}

func GetData() *map[string]string {
	return &data
}
func Delete(k string) {
	mu.Lock()
	defer mu.Unlock()
	delete(data, k)
}
func Set(k, v string) {
	mu.Lock()
	defer mu.Unlock()
	data[k] = v
}
func Get(k string) string {
	mu.Lock()
	defer mu.Unlock()
	return data[k]
}
func Exists(k string) bool {
	mu.Lock()
	defer mu.Unlock()
	_, ok := data[k]
	return ok
}
func Keys() []string {
	mu.Lock()
	defer mu.Unlock()
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
func Replace(newkv map[string]string) {
	mu.Lock()
	defer mu.Unlock()
	// take old reference and garbage collect memory
	data = newkv
}
