// Key value store for global configuration data

package data

import (
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
func Keys() []string {
	mu.Lock()
	defer mu.Unlock()
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}
