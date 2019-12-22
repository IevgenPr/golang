package main

import (
	"testing"
)

func TestMaps(t *testing.T) {

	t.Run("Test maps",
		func(t *testing.T) {
			want := map[string]int{"a": 1}
			got := printMap(&want)

			if got == &want {
				t.Errorf("printMap want/got \n%w \n%w", want, got)
			}

		})
}
