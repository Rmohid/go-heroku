// Template for command line application using JSON over http

package main

import (
	"github.com/rmohid/go-template/config"
	"github.com/rmohid/go-template/webExternal"
	"github.com/rmohid/go-template/webInternal"
	"os"
)

var err error

func main() {

	if err = config.ParseArgs(); err != nil {
		os.Exit(1)
	}
	go webInternal.Run()
	webExternal.Run()
}
