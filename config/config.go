// Key value store for global configuration data

package config

import (
	"flag"
	"github.com/rmohid/go-template/config/data"
	"github.com/rmohid/go-template/config/webInternal"
)

type Option struct {
	Value                      *string
	Name, Default, Description string
}

var (
	indexed map[string]int
	options []Option
)

func init() {
	indexed = make(map[string]int)
	opts := [][]string{
		{"readableJson", "yes", "pretty print json output"},
	}

	PushArgs(opts)
}
func Delete(k string) {
	data.Delete(k)
}
func Set(k, v string) {
	data.Set(k, v)
}
func Get(k string) string {
	return data.Get(k)
}
func Exists(k string) bool {
	return data.Exists(k)
}
func Keys() []string {
	return data.Keys()
}
func Replace(newkv map[string]string) {
	data.Replace(newkv)
}
func PushArgs(inOpts [][]string) error {
	for i, _ := range inOpts {
		Name, Default, Description := inOpts[i][0], inOpts[i][1], inOpts[i][2]
		j, exists := indexed[Name]
		if exists {
			options[j].Name, options[j].Default, options[j].Description = Name, Default, Description
			continue
		}
		indexed[Name] = i
		options = append(options, Option{nil, Name, Default, Description})
	}
	return nil
}
func ParseArgs(inOpts [][]string) error {

	PushArgs(inOpts)
	for i, _ := range options {
		var elem = &options[i]
		elem.Value = flag.String(elem.Name, elem.Default, elem.Description)
	}
	// nothing is actally done until parse is called
	flag.Parse()
	for _, elem := range options {
		data.Set(elem.Name, *elem.Value)
	}

	// Start the internal admin web interface
	go webInternal.Run()
	return nil
}
