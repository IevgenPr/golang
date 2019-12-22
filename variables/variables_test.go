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
			if got != "0" {
				t.Errorf("printVar expects 0 = %w", got)
			}
		})
	t.Run("Test another function",
		func(t *testing.T) {
			got := changeName("blah")
			want := "blah, what a name!"
			if got != want {
				t.Errorf("printVar want/got \n%w \n%w", want, got)
			}
		})
}
