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
	search <query>    Search by package name and description`)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) == 0 {
		usage()
	}

	cmd := args[0]
	args = args[1:]

	runCommand(cmd, args)
}
