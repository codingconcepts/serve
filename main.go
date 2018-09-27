package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)

	dir := flag.String("d", "", "relative or absolute path of directory to serve")
	port := flag.Int("p", 8080, "service port")
	flag.Parse()

	if *dir == "" || *port == 0 {
		flag.Usage()
		os.Exit(1)
	}

	_, err := os.Stat(*dir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("directory %q does not exist", *dir)
		}
		log.Fatalf("can't share %q: %v", *dir, err)
	}

	if *dir, err = filepath.Abs(*dir); err != nil {
		log.Fatalf("unable to get absolute path for directory %q", *dir)
	}

	log.Printf("Serving from: %s\nListening on: http://localhost:%d", *dir, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(*dir))))
}
