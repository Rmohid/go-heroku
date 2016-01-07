// Template for console based debug tracing

package dbg

import (
	"fmt"
	"github.com/rmohid/go-heroku/Godeps/_workspace/src/github.com/rmohid/go-heroku/config"
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
	writers["file"] = filewriter()

	// define all default options
	var opts = [][]string{
		{"dbg.debugWriter", "stderr", "debug log output sink"},
		{"dbg.verbosity", "0", "verbosity level for debug output"},
		{"dbg.logfile", "config.log", "filename for log collection"},
		{"dbg.httpUrl", "", "http server for log delivery"},
	}

	config.PushArgs(opts)
}
func ErrLog(verbosity int, format string, a ...interface{}) {
	val, err := strconv.Atoi(config.Get("dbg.verbosity"))
	if err != nil {
		log.Fatal("dbg.ErrLog:", err)
	}
	if val >= verbosity {
		fmt.Fprintf(os.Stderr, format, a...)
	}
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
	str = fmt.Sprintf("http://%s?%s=%s", str, time.Now().Format("2006-01-02T15:04:05.999Z07:00"), url.QueryEscape(payload))
	resp, err := http.Get(str)
	if err != nil {
		return 0, fmt.Errorf("dbg.httpWriter:", err)
	}
	defer resp.Body.Close()
	return 0, nil
}

func filewriter() io.Writer {
	var h fileWriter
	return h
}

type fileWriter struct {
}

func (h fileWriter) Write(p []byte) (n int, err error) {
	str := config.Get("dbg.logfile")
	if str == "" {
		return 0, nil
	}
	f, err := os.OpenFile(str, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("dbg.filewriter:", err)
	}
	defer f.Close()
	return f.Write(p)
}
