package main

import (
	"flag"
	"fmt"
	"goeflo/grmn-server/rest"
	"log"
	"os"
)

func main() {
	fmt.Println("grmn-server ...")

	activitiesDir := flag.String("directory", "activities", "fit file activities directory")
	verbose := flag.Bool("verbose", false, "more verbose output")
	port := flag.Int("port", 8088, "server listen on this port")
	flag.Parse()

	if *verbose {
		fmt.Printf("starting with activitiesDir=%v, port=%v, verbose=%v\n", *activitiesDir, *port, *verbose)
	}

	if !pathExists(*activitiesDir) {
		log.Fatalf("path to activities '%v' does not exists, please check given directory parameter", *activitiesDir)
	}

	rest.Start(rest.RestOpts{Activities: *activitiesDir, Port: *port, Verbose: *verbose})
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}
