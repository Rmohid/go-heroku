// Template for console based debug tracing

package dbg

import (
	"fmt"
	"github.com/rmohid/go-template/config"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	writers = map[string](io.Writer){}
)

func init() {
	writers[""] = devnull()
	writers["devnull"] = devnull()
	writers["stdout"] = os.Stdout
	writers["stderr"] = os.Stderr

	// define all default options
	var opts = [][]string{
		{"dbg.debugWriter", "stderr", "default debug log output"},
		{"dbg.verbosity", "0", "verbosity level for debug output"},
	}

	config.PushArgs(opts)
}

func devnull() io.Writer {
	null, err := os.Open(os.DevNull)
	if err != nil {
		log.Fatal("dbg.devnull:", err)
	}
	return null
}

func Log(verbosity int, a ...interface{}) {
	val, err := strconv.Atoi(config.Get("dbg.verbosity"))
	if err != nil {
		log.Fatal("dbg.Log():", err)
	}
	if val >= verbosity {
		fmt.Fprintln(writers[config.Get("dbg.debugWriter")],
			a...)
	}
}
