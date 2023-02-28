package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"urlshort"
	"fmt"

)

func main() {
	filetype, filename := parseArg()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	filedata , err := parseData(filename)
	if err != nil {
		log.Fatal(err)
	}

	// var handler http.HandlerFunc
	handler := func() http.HandlerFunc {
		if filetype == "yaml" || filetype == "yml" {
			yamlHandler, err := urlshort.YAMLHandler([]byte(filedata), mapHandler)
			if err != nil {
				panic(err)
			}
			return yamlHandler
		}else {
			JSONHandler, err := urlshort.JSONHandler([]byte(filedata), mapHandler)
			if err != nil {
				panic(err)
			}
			return JSONHandler
		}
	}


	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler())
}


func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func parseData(filename string)([]byte, error){
	return ioutil.ReadFile(filename)
} 

func parseArg()(string, string){
	filetype := flag.String("filetype", "yml", "File type. Example: yaml, json, etc.")
	filename := flag.String("filename", "", "Yaml paths file")
	flag.Parse()
	return *filetype, *filename
}