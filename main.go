package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: saur <command> [...]
	
Commands:
	search <search term(s)>    Search by package name and description
	get    <package(s)>        Get PKGBUILD`)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	if err := runCommand(args[0], args[1:]); err != nil {
		log.Fatal(err)
	}
}
