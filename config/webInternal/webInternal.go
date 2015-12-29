// Handler for administrative web interface

package webInternal

import (
	"encoding/json"
	"fmt"
	"github.com/rmohid/go-template/config/data"
	"log"
	"net/http"
	"strings"
)

func Run() {
	serverInternal := http.NewServeMux()
	serverInternal.HandleFunc("/", handler)
	serverInternal.HandleFunc("/json", handleJson)
	log.Fatal(http.ListenAndServe(data.Get("portInternal"), serverInternal))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
			data.Set(k, strings.Join(v, " "))
		}
	} else {
		configKeys := data.Keys()
		for _, k := range configKeys {
			fmt.Fprintf(w, "Config[%q] = %q\n", k, data.Get(k))
		}
	}
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	dat, err := json.MarshalIndent(data.GetData(), "", "  ")
	if err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, "%s", dat)
}
