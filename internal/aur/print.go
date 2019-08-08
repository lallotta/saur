package aur

import (
	"fmt"
	"time"

	"github.com/lallotta/saur/internal/colors"
)

// TODO: implement functions to print package info

// Pkgs holds the packages returned from a query
type Pkgs []Package

// Print prints search results
func (pkgs Pkgs) Print() {
	for _, pkg := range pkgs {
		fmt.Printf(colors.Bold("%s %s"), pkg.Name, colors.Blue(pkg.Version))

		if pkg.OutOfDate > 0 {
			year, month, day := time.Unix(int64(pkg.OutOfDate), 0).Date()
			str := fmt.Sprintf(" Out of Date %d-%02d-%02d", year, month, day)
			fmt.Print(colors.Bold(colors.Red(str)))
		}

		fmt.Printf("\n    %s\n", pkg.Description)
	}
}
