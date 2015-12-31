// Key value web api for configuration data
// See github.com/rmohid/go-template for detailed description

package config

import (
	"flag"
	"fmt"
	"github.com/rmohid/go-template/config/data"
	"github.com/rmohid/go-template/config/webInternal"
)

type Option struct {
	Value                      *string
	Name, Default, Description string
}

const (
	NameIdx = iota
	DefaultIdx
	DescriptionIdx
)

var (
	indexed map[string]int
	options []Option
)

func init() {
	indexed = make(map[string]int)

	// default options for config package
	opts := [][]string{
		{"config.readableJson", "yes", "pretty print json output"},
		{"config.enableFlagParse", "yes", "allow config to flag.Parse()"},
		{"config.silentWebPrompt", "no", "display internal port used"},
		{"config.portInternal", "localhost:7100", "internal web port"},
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
		Name, Default, Description := inOpts[i][NameIdx], inOpts[i][DefaultIdx], inOpts[i][DescriptionIdx]
		j, exists := indexed[Name]
		if exists {
			options[j].Name, options[j].Default, options[j].Description = Name, Default, Description
			continue
		}
		indexed[Name] = i
		options = append(options, Option{nil, Name, Default, Description})
		data.Set(Name, Default)
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
	if Get("config.enableFlagParse") == "yes" {
		flag.Parse()
	}
	for _, elem := range options {
		data.Set(elem.Name, *elem.Value)
	}

	// Start the internal admin web interface
	if Get("silentWebPrompt") == "no" {
		fmt.Println("listening on", Get("portInternal"))
	}
	go webInternal.Run()
	return nil
}
