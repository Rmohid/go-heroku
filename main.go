// Template for command line application using JSON over http

package main

import (
	"fmt"
	"github.com/rmohid/go-heroku/Godeps/_workspace/src/github.com/rmohid/go-heroku/config"
	"github.com/rmohid/go-heroku/Godeps/_workspace/src/github.com/rmohid/go-heroku/dbg"
	"github.com/rmohid/go-heroku/Godeps/_workspace/src/github.com/rmohid/go-heroku/webExternal"
	"os"
	"time"
)

var err error

func main() {

	// define all string based options
	var opts = [][]string{
		{"portExternal", "localhost:7000", "external web port"},
		{"config.portInternal", "localhost:7100"},
		{"dbg.httpUrl", "localhost:7000"},
		{"dbg.verbosity", "1"},
	}

	if err = config.ParseArgs(opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	dbg.Log(2, config.Dump())
	dbg.Log(0, "Starting..")
	fmt.Println("listening on", config.Get("portExternal"))

	go test()
	webExternal.Run()
}

func test() {
	for {
		time.Sleep(1 * time.Second)
		dbg.Log(1, "Debug log 1 ", time.Now().Format(time.StampMilli))
		dbg.Log(2, "Debug log 2 ", time.Now().Format(time.Kitchen))
	}
}
