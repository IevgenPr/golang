package main

import "testing"

func TestVariables(t *testing.T) {
	// boring test assertion
	if module != "golang module var" {
		t.Errorf("got: \n%w ", module)
	}

	// sowewhat more interesting test
	want := [2]interface{}{"golang module var", 10}
	if [2]interface{}{module, intnum} != want {
		t.Errorf("want/got: \n%w \n%w", want, [2]interface{}{module, intnum})
	}
	// nested test case - pretty neat
	t.Run("Test one function",
		func(t *testing.T) {
			got := printVar(0)
			if got != "00" {
				t.Errorf("printVar expects 0 = %w", got)
			}
		})
	t.Run("Test another function",
		func(t *testing.T) {
			got := changeName("blah")
			if got != "blah" {
				t.Errorf("printVar expects blah = %w", got)
			}
		})
}
