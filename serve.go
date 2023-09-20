package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	version string
)

func main() {
	log.SetFlags(0)

	dir := flag.String("d", "", "relative or absolute path of directory to serve")
	port := flag.Int("p", 8080, "service port")
	versionFlag := flag.Bool("version", false, "display the current version number")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		return
	}

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

	// Doing this for users sharing on a local network.
	addr := localAddr()
	log.Printf("Serving from: %s\nListening on: %s:%d", *dir, addr, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(*dir))))
}

func localAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "localhost"
}
