package aur

import (
	"fmt"
	"strings"
	"time"

	"github.com/lallotta/saur/internal/colors"
)

// Pkgs holds the packages returned from a query
type Pkgs []Package

// Print prints search results
func (pkgs Pkgs) Print() {
	for _, pkg := range pkgs {
		fmt.Print(colors.Bold(pkg.Name + " " + colors.Blue(pkg.Version)))

		if pkg.OutOfDate > 0 {
			str := fmt.Sprintf(" (Out of Date %s)", formatDate(pkg.OutOfDate))
			fmt.Print(colors.Bold(colors.Red(str)))
		}

		if pkg.Maintainer == "" {
			fmt.Print(colors.Bold(colors.Red(" (Orphaned)")))
		}

		fmt.Println("\n    " + pkg.Description)
	}
}

// PrintInfo prints package information
func (pkg Package) PrintInfo() {
	printFieldInfo("Name", pkg.Name)
	printFieldInfo("Version", pkg.Version)
	printFieldInfo("Description", pkg.Description)
	printFieldInfo("URL", pkg.URL)
	printFieldInfo("Keywords", strings.Join(pkg.Keywords, " "))
	printFieldInfo("Licenses", strings.Join(pkg.License, " "))
	printFieldInfo("Groups", strings.Join(pkg.Groups, " "))
	printFieldInfo("Provides", strings.Join(pkg.Provides, " "))
	printFieldInfo("Depends On", strings.Join(pkg.Depends, " "))
	printFieldInfo("Optional Deps", strings.Join(pkg.OptDepends, " "))
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
