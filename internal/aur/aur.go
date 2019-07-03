package aur

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
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

type queryResult struct {
	Type    string
	Error   string
	Results PkgList
}

func request(v url.Values) (PkgList, error) {
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
func Search(query string) (PkgList, error) {
	v := url.Values{}
	v.Set("type", "search")
	v.Set("arg", query)
	return request(v)
}

// Info requests package information for the given package names
func Info(pkgNames []string) (PkgList, error) {
	v := url.Values{}
	v.Set("type", "info")

	for _, name := range pkgNames {
		v.Add("arg[]", name)
	}

	return request(v)
}
