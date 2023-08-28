package main

import (
	"gceny"
	"net/http"
)

// test example main func
func main() {

	engine := gceny.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	_ = http.ListenAndServe(":9090", engine)
}
