// Template for command line application using JSON over http

package config

import (
	"flag"
)

var D map[string]string

// Default values for the application
func init() {
	D = make(map[string]string)
}

func parseArgs() error {
	D["portExternal"] = *flag.String("portExternal", "localhost:8000", "external web port")
	D["portInternal"] = *flag.String("portInternal", "localhost:8100", "Internal web port")
	return nil
}
