package main

import (
	"fmt"
	"log"
)

func runCommand(cmd string, args []string) {
	switch cmd {
	case "search":
		runSearch(args)
	default:
		log.Fatalf("invalid command: '%s' (use -h for help)", cmd)
	}
}

func runSearch(targets []string) {
	if len(targets) == 0 {
		log.Fatal("saur search: no targets specified")
	}

	fmt.Println("search command")
	fmt.Println("targets:", targets)
}
