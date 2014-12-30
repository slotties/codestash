package main

import (
//    "fmt"
    "net/http"
    "flag"
    "path"
)

var directory string

func main() {
	port := flag.String("p", "8080", "the port to bind on (ports below 1024 require root permissions)")
	directoryPath := flag.String("d", ".", "directory to serve")
	flag.Parse()

	directory = *directoryPath

    http.HandleFunc("/", serveFile)
    http.ListenAndServe(":" + *port, nil)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	fullFileName := path.Join(directory, r.URL.Path[1:])
	http.ServeFile(w, r, fullFileName)
}