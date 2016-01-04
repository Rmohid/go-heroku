// Template for console based debug tracing

package dbg

import (
	"fmt"
	"github.com/rmohid/go-template/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	writers = map[string](io.Writer){}
)

func init() {
	writers[""] = devnull()
	writers["devnull"] = devnull()
	writers["stdout"] = os.Stdout
	writers["stderr"] = os.Stderr
	writers["http"] = httpwriter()

	// define all default options
	var opts = [][]string{
		{"dbg.debugWriter", "stderr", "default debug log output"},
		{"dbg.verbosity", "0", "verbosity level for debug output"},
		{"dbg.httpUrl", "", "http server for log delivery"},
	}

	config.PushArgs(opts)
}
func Log(verbosity int, a ...interface{}) {
	val, err := strconv.Atoi(config.Get("dbg.verbosity"))
	if err != nil {
		log.Fatal("dbg.Log:", err)
	}
	if val >= verbosity {
		fmt.Fprintln(writers[config.Get("dbg.debugWriter")],
			a...)
	}
}
func devnull() io.Writer {
	null, err := os.Open(os.DevNull)
	if err != nil {
		log.Fatal("dbg.devnull:", err)
	}
	return null
}

func httpwriter() io.Writer {
	var h httpWriter
	return h
}

type httpWriter struct {
}

func (h httpWriter) Write(p []byte) (n int, err error) {
	str := config.Get("dbg.httpUrl")
	if str == "" {
		return 0, nil
	}
	payload := fmt.Sprintf(string(p[:]))
	resp, err := http.PostForm(str, url.Values{time.Now().Format("15:04:05 "): {payload}})
	if err != nil {
		return 0, fmt.Errorf("dbg.httpWriter:", err)
	}
	defer resp.Body.Close()
	return 0, nil
}