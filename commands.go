package main

import (
	"log"
	"strings"

	"github.com/lallotta/saur/internal/aur"
)

func runCommand(cmd string, args []string) {
	switch cmd {
	case "search":
		runSearch(args)
	default:
		log.Fatalf("invalid command: '%s' (use -h for help)", cmd)
	}
}

func runSearch(terms []string) {
	if len(terms) == 0 {
		log.Fatal("saur search: no search terms specified")
	}

	// TODO: search should be performed on a single term
	// and then filtered based on unused terms to find matches.
	// Joining terms together in this way may not return all expected results.
	query := strings.Join(terms, " ")

	results, err := aur.Search(query)
	if err != nil {
		log.Fatalln("AUR query returned an error:", err)
	}

	results.Print()
}
