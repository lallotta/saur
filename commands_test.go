package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestInvalidCommand(t *testing.T) {
	input := "foo"
	err := runCommand(input, []string{})
	want := fmt.Sprintf("invalid command: '%s'", input)

	if err == nil {
		t.Errorf("'%s' ran successfully", input)
	} else {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("%s: got %q; want %q", input, err, want)
		}
	}
}

func TestBadCommandLine(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"search", "no search terms specified"},
		{"info", "no packages specified"},
		{"get", "no targets specified"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			err := runCommand(test.in, []string{})
			if err == nil {
				t.Errorf("'%s' ran successfully with no arguments", test.in)
			} else {
				if !strings.Contains(err.Error(), test.want) {
					t.Errorf("%s: got %q; want %q", test.in, err, test.want)
				}
			}
		})
	}
}
