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
