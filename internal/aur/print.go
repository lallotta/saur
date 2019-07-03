package aur

import (
	"fmt"
	"time"
)

// TODO: colorize output and implement functions print package info

// PkgList holds the packages returned from a query
type PkgList []Package

// Print prints search results
func (pkgs PkgList) Print() {
	for _, pkg := range pkgs {
		fmt.Printf("%s %s", pkg.Name, pkg.Version)

		if pkg.OutOfDate > 0 {
			year, month, day := time.Unix(int64(pkg.OutOfDate), 0).Date()
			fmt.Printf(" (Out of Date %d-%02d-%d)", year, month, day)
		}

		fmt.Printf("\n    %s\n", pkg.Description)
	}
}
