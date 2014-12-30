package main

import (
    "net/http"
    "flag"
)

func main() {
	port := flag.String("p", "8080", "the port to bind on (ports below 1024 require root permissions)")
	directoryPath := flag.String("d", ".", "directory to serve")
	flag.Parse()

    http.ListenAndServe(":" + *port, http.FileServer(http.Dir(*directoryPath)))
}