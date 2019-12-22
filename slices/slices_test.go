package main

import (
	"fmt"
	"testing"
)

func TestVariables(t *testing.T) {

	t.Run("Lazy",
		func(t *testing.T) {
			fmt.Println("No sense, but passes")
		})
}
