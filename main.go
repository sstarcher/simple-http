package main

import (
	"flag"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var debug = flag.Bool("debug", false, "enables debug logging")

func init() {
	flag.Parse()
	if debug != nil && *debug {
		log.SetLevel(log.DebugLevel)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("http://%s%s", r.Host, r.URL.Path)
	if r.Method == http.MethodGet && (r.URL.Path == "/" || r.URL.Path == "") {
		acceptHeader := r.Header.Get("Accept")
		if strings.Contains(acceptHeader, "application/json") {
			w.Write([]byte(`{"message": "Good morning"}`))
		} else {
			w.Write([]byte("<p>Hello, World</p>"))
		}
	}
}
