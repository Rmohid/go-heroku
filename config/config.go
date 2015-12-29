// Key value store for global configuration data

package config

import (
	"flag"
	"github.com/rmohid/go-template/config/data"
	"github.com/rmohid/go-template/config/webInternal"
)

type Option struct {
	Value                     *string
	Name, Default, Descripton string
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
func Keys() []string {
	return data.Keys()
}

func ParseArgs(inOpts [][]string) error {

	var opts = []Option{}
	for i, _ := range inOpts {
		opts = append(opts, Option{nil, inOpts[i][0], inOpts[i][1], inOpts[i][2]})
		var elem = &opts[i]
		elem.Value = flag.String(elem.Name, elem.Default, elem.Descripton)
	}
	flag.Parse()
	for _, elem := range opts {
		data.Set(elem.Name, *elem.Value)
	}

	// Start the internal admin web interface
	go webInternal.Run()
	return nil
}
