package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lallotta/saur/internal/aur"
	"github.com/lallotta/saur/internal/aur/fetch"
	"github.com/lallotta/saur/internal/colors"
)

var (
	infoPrefix    string
	warningPrefix string
	errorPrefix   string
)

func init() {
	prefix := "::"
	infoPrefix = colors.Cyan(prefix)
	warningPrefix = colors.Yellow(prefix)
	errorPrefix = colors.Red(prefix)
}

type errorSlice struct {
	errors []error
}

func (e *errorSlice) add(err error) {
	e.errors = append(e.errors, err)
}

func (e *errorSlice) Error() string {
	s := ""
	for _, err := range e.errors {
		s += err.Error() + "\n"
	}
	return s
}

func (e *errorSlice) returnError() error {
	if len(e.errors) > 0 {
		return e
	}
	return nil
}

func runCommand(cmd string, args []string) error {
	switch cmd {
	case "search":
		return runSearch(args)
	case "get":
		return runGet(args)
	}
	return fmt.Errorf("%s invalid command: '%s' (use -h for help)", errorPrefix, cmd)
}

func runSearch(terms []string) error {
	if len(terms) == 0 {
		return errors.New(errorPrefix + "no search terms specified")
	}

	// TODO: search should be performed on a single term
	// and then filtered based on unused terms to find matches.
	// Joining terms together in this way may not return all expected results.
	query := strings.Join(terms, " ")

	results, err := aur.Search(query)
	if err != nil {
		return fmt.Errorf("%s error querying AUR: %v", errorPrefix, err)
	}

	results.Print()

	return nil
}

func runGet(targets []string) error {
	if len(targets) == 0 {
		return errors.New(errorPrefix + "no packages specified")
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs errorSlice

	pkgs, err := aur.Info(targets)
	if err != nil {
		return fmt.Errorf("%s error querying AUR: %v", errorPrefix, err)
	}

	// Create map of package names
	found := make(map[string]bool)
	for _, pkg := range pkgs {
		found[pkg.Name] = true
	}

	download := func(name string) {
		defer wg.Done()
		if err := fetch.GetPkgbuild(name); err != nil {
			mu.Lock()
			errs.add(fmt.Errorf("%s error downloading '%s': %v", errorPrefix, name, err))
			mu.Unlock()
			return
		}
		log.Println("Finished downloading", name)
	}

	log.Println(colors.Bold(infoPrefix + " Downloading PKGBUILDs..."))
	for _, target := range targets {
		if !found[target] {
			errs.add(fmt.Errorf("%s target not found: '%s'", errorPrefix, target))
			continue
		}

		_, err = os.Stat(filepath.Join(dir, target))
		if err == nil {
			log.Printf("%s '%s' already exists", warningPrefix, target)
			continue
		} else if !os.IsNotExist(err) {
			log.Println(errorPrefix, err)
			continue
		}
		wg.Add(1)
		go download(target)

	}

	wg.Wait()

	return errs.returnError()
}
