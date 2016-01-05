package webExternal

import (
	"log"
	"os"
	"testing"
)

var (
	Trace *log.Logger
)

func init() {
	// for Trace.Println(Keys())
	Trace = log.New(os.Stdout, "**: ", log.Lshortfile)
}

func TestStub(t *testing.T) {
	var expected, got string
	if expected != got {
		t.Error(
			"Expected", expected,
			"got", got,
		)
	}
}
