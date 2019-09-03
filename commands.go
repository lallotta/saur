package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/lallotta/saur/internal/aur"
	"github.com/lallotta/saur/internal/aur/fetch"
	"github.com/lallotta/saur/internal/colors"
)

const arrow = "==>"

// errorGroup is a collection of errors
type errorGroup struct {
	buf bytes.Buffer
}

// add adds a new error to the group
func (e *errorGroup) add(err error) {
	e.buf.WriteString(err.Error() + "\n")
}

func (e *errorGroup) Error() string {
	return e.buf.String()
}

// returnError returns an error representing
// any collected errors
func (e *errorGroup) returnError() error {
	if e.buf.Len() > 0 {
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
	case "info":
		return runInfo(args)
	}
	return errorf("invalid command: '%s' (use -h for help)", cmd)
}

func runSearch(terms []string) error {
	if len(terms) == 0 {
		return errorf("no search terms specified")
	}

	// TODO: Perform search on a single term and filter
	// based on unused terms to find matches. Joining terms
	// together in this way may not return all expected results.
	query := strings.Join(terms, " ")

	results, err := aur.Search(query)
	if err != nil {
		return errorf("AUR query error: %v", err)
	}

	results.Print()

	return nil
}

func runGet(targets []string) error {
	if len(targets) == 0 {
		return errorf("no targets specified")
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs errorGroup

	colonPrintln("Requesting package info...")
	pkgs, err := aur.Info(targets)
	if err != nil {
		return errorf("AUR query error: %v", err)
	}

	found := make(map[string]struct{})
	for _, pkg := range pkgs {
		found[pkg.Name] = struct{}{}
	}

	download := func(name string) {
		defer wg.Done()
		if err := fetch.GetPkgbuild(name); err != nil {
			mu.Lock()
			errs.add(errorf("failure downloading '%s': %v", name, err))
			mu.Unlock()
			return
		}
		fmt.Println("Downloaded PKGBUILD:", name)
	}

	colonPrintln("Downloading PKGBUILDs...")
	for _, target := range targets {
		if _, ok := found[target]; !ok {
			errs.add(errorf("package not found: '%s'", target))
		} else {
			_, err = os.Stat(target)
			if err == nil {
				warnPrintf("'%s' already exists\n", target)
				continue
			} else if !os.IsNotExist(err) {
				log.Println(colors.Bold(colors.Red(arrow)), err)
				continue
			}

			wg.Add(1)
			go download(target)
		}
	}

	wg.Wait()

	return errs.returnError()
}

func runInfo(targets []string) error {
	if len(targets) == 0 {
		return errorf("no packages specified")
	}

	var errs errorGroup

	pkgs, err := aur.Info(targets)
	if err != nil {
		return errorf("AUR query error: %v", err)
	}

	found := make(map[string]int)
	for i, pkg := range pkgs {
		found[pkg.Name] = i
	}

	for _, target := range targets {
		if idx, ok := found[target]; !ok {
			errs.add(errorf("package not found: '%s'", target))
		} else {
			pkg := pkgs[idx]
			pkg.PrintInfo()
		}
	}

	return errs.returnError()
}

func errorf(format string, a ...interface{}) error {
	format = colors.Bold(colors.Red("error: ")) + format
	return fmt.Errorf(format, a...)
}

func colonPrintln(a ...interface{}) (n int, err error) {
	colon := colors.Bold(colors.Cyan("::"))
	return fmt.Println(colon, colors.Bold(fmt.Sprint(a...)))
}

func warnPrintf(format string, a ...interface{}) (n int, err error) {
	format = colors.Bold(colors.Yellow(arrow)) + " " + format
	return fmt.Printf(format, a...)
}
