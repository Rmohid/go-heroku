// Key value store for global configuration data, replace with whatever backend you like

package data

import (
	"sort"
	"sync"
)

type kvData struct {
	key, value string
	exists     bool
}

var (
	data   map[string]string
	mu     sync.Mutex
	get    = make(chan string)
	remove = make(chan string)
	keys   = make(chan []string)
	set    = make(chan kvData)
	exists = make(chan kvData)
)

func init() {
	data = make(map[string]string)
	go broker()
}

func broker() {
	for {
		select {
		case key := <-get:
			get <- data[key]
		case kv := <-set:
			data[kv.key] = kv.value
		case k := <-remove:
			delete(data, k)
		case kv := <-exists:
			_, ok := data[kv.key]
			exists <- kvData{"", "", ok}
		case <-keys:
			list := make([]string, 0, len(data))
			for k := range data {
				list = append(list, k)
			}
			sort.Strings(list)
			keys <- list
		}
	}
}
func GetData() *map[string]string {
	return &data
}
func Delete(k string) {
	remove <- k
}
func Set(k, v string) {
	if k == "" {
		return
	}
	set <- kvData{k, v, true}
}
func Get(k string) string {
	get <- k
	return <-get
}
func Exists(k string) bool {
	exists <- kvData{k, "", true}
	return (<-exists).exists
}
func Keys() []string {
	keys <- []string{}
	return <-keys
}
func Replace(newkv map[string]string) {
	mu.Lock()
	defer mu.Unlock()
	// take old reference and garbage collect memory
	data = newkv
}
