// Template for command line application using JSON over http

package main

import (
	"fmt"
	"github.com/rmohid/go-template/config"
	"github.com/rmohid/go-template/dbg"
	"github.com/rmohid/go-template/webExternal"
	"os"
)

var err error

func main() {

	// define all string based options
	var opts = [][]string{
		{"config.portInternal", "localhost:7100", "external web port"},
		{"portExternal", "localhost:7000", "external web port"},
	}

	if err = config.ParseArgs(opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	dbg.Log(0, "Starting..")
	fmt.Println("listening on", config.Get("portExternal"))

	webExternal.Run()
}
