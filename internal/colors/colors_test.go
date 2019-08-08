package colors

import (
	"fmt"
	"testing"
)

func TestColors(t *testing.T) {
	tests := []struct {
		name  string
		param parameter
		fn    func(string) string
	}{
		{"Bold", bold, Bold},
		{"Red", redFg, Red},
		{"Green", greenFg, Green},
		{"Brown", brownFg, Brown},
		{"Blue", blueFg, Blue},
		{"Magenta", magentaFg, Magenta},
		{"Cyan", cyanFg, Cyan},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := fmt.Sprintf("\x1b[%dm%s\x1b[0m", test.param, test.name)

			// Visual test
			fmt.Println(want)

			got := test.fn(test.name)
			if got != want {
				t.Errorf("%s: got %q; want %q", test.name, got, want)
			}
		})
	}
}
