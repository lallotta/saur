package main

import (
	"strings"
	"testing"
)

func TestBadCommandLine(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"search", "no search terms specified"},
		{"info", "no packages specified"},
		{"get", "no targets specified"},
		{"foo", "invalid command: 'foo'"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			err := runCommand(test.in, []string{})
			if err == nil {
				t.Errorf("'%s' ran successfully", test.in)
			} else if !strings.Contains(err.Error(), test.want) {
				t.Errorf("%s: got %q; want %q", test.in, err, test.want)
			}
		})
	}
}
