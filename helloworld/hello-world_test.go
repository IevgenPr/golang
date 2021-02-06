package main

import "testing"

func TestHelloWorld(t *testing.T) {
	name = "helo"

	if name != "helo" {
		t.Errorf("got: \n%s ", name)
	}
	want := [3]interface{}{"tests", "test2", 10.0}
	if [3]interface{}{tests, test2, fl} != want {
		t.Errorf("want/got: \n%s \n%s", want, [3]interface{}{tests, test2, fl})
	}
}
