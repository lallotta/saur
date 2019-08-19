package aur

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lallotta/saur/internal/colors"
)

// Package contains package info
type Package struct {
	ID             int
	Name           string
	PackageBaseID  int
	PackageBase    string
	Version        string
	Description    string
	URL            string
	NumVotes       int
	Popularity     float64
	OutOfDate      int
	Maintainer     string
	FirstSubmitted int
	LastModified   int
	URLPath        string
	Depends        []string
	MakeDepends    []string
	OptDepends     []string
	CheckDepends   []string
	Conflicts      []string
	Provides       []string
	Replaces       []string
	Groups         []string
	License        []string
	Keywords       []string
}

// Pkgs holds the packages returned from a query
type Pkgs []Package

type queryResult struct {
	Type    string
	Error   string
	Results Pkgs
}

func request(v url.Values) (Pkgs, error) {
	rpcURL := "https://aur.archlinux.org/rpc/?"
	v.Set("v", "5")

	resp, err := http.Get(rpcURL + v.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := new(queryResult)

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}

	if result.Type == "error" {
		return nil, errors.New(result.Error)
	}

	return result.Results, nil
}

// Search performs a search for packages by the default name-desc field
func Search(query string) (Pkgs, error) {
	v := url.Values{}
	v.Set("type", "search")
	v.Set("arg", query)
	return request(v)
}

// Info requests package information for the given package names
func Info(pkgNames []string) (Pkgs, error) {
	v := url.Values{}
	v.Set("type", "info")

	for _, name := range pkgNames {
		v.Add("arg[]", name)
	}

	return request(v)
}

// Print prints search results
func (pkgs Pkgs) Print() {
	for _, pkg := range pkgs {
		var buf bytes.Buffer

		buf.WriteString(colors.Bold(pkg.Name + " " + colors.Blue(pkg.Version)))

		if pkg.OutOfDate > 0 {
			str := fmt.Sprintf(" (Out of Date %s)", formatDate(pkg.OutOfDate))
			buf.WriteString(colors.Bold(colors.Red(str)))
		}

		if pkg.Maintainer == "" {
			buf.WriteString(colors.Bold(colors.Red(" (Orphaned)")))
		}

		buf.WriteString("\n    " + pkg.Description)

		fmt.Println(&buf)
	}
}

// PrintInfo prints package information
func (pkg Package) PrintInfo() {
	printFieldInfo("Name", pkg.Name)
	printFieldInfo("Version", pkg.Version)
	printFieldInfo("Description", pkg.Description)
	printFieldInfo("URL", pkg.URL)
	printFieldInfo("Licenses", strings.Join(pkg.License, " "))
	printFieldInfo("Groups", strings.Join(pkg.Groups, " "))
	printFieldInfo("Depends On", strings.Join(pkg.Depends, " "))
	printFieldInfo("Make Deps", strings.Join(pkg.MakeDepends, " "))
	printFieldInfo("Optional Deps", strings.Join(pkg.OptDepends, " "))
	printFieldInfo("Check Deps", strings.Join(pkg.CheckDepends, " "))
	printFieldInfo("Conflicts With", strings.Join(pkg.Conflicts, " "))
	printFieldInfo("Replaces", strings.Join(pkg.Replaces, " "))
	printFieldInfo("Maintainer", pkg.Maintainer)
	printFieldInfo("Votes", fmt.Sprintf("%d", pkg.NumVotes))
	printFieldInfo("Popularity", fmt.Sprintf("%f", pkg.Popularity))
	printFieldInfo("First Submitted", formatDateTime(pkg.FirstSubmitted))
	printFieldInfo("Last Modified", formatDateTime(pkg.LastModified))

	if pkg.OutOfDate > 0 {
		printFieldInfo("Out of Date", formatDate(pkg.OutOfDate))
	} else {
		printFieldInfo("Out of Date", "No")
	}

	fmt.Println()
}

func printFieldInfo(field, value string) {
	if value == "" {
		value = "None"
	}
	fmt.Printf(colors.Bold("%-15s :")+" %s\n", field, value)
}

func formatDate(t int) string {
	return time.Unix(int64(t), 0).Format("2006-01-02")
}

func formatDateTime(t int) string {
	tm := time.Unix(int64(t), 0)
	return tm.Format("Mon 02 Jan 2006 03:04:05 PM MST")
}
