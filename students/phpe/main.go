package main

import (
	"fmt"
	"net/http"

	// "github.com/gophercises/urlshort"
)

// func main() {
// 	mux := defaultMux()

// 	// Build the MapHandler using the mux as the fallback
// 	pathsToUrls := map[string]string{
// 		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
// 		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
// 	}
// 	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

// 	// Build the YAMLHandler using the mapHandler as the
// 	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `
// 	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Starting the server on :8080")
// 	http.ListenAndServe(":8080", yamlHandler)
// }

// func defaultMux() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	return mux
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, world!")
// }

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}