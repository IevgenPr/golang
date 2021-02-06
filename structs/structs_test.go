package main

import "testing"

func TestVariables(t *testing.T) {

	t.Run("Test another function",
		func(t *testing.T) {

			want := courseMeta{"Nice Fellow", "Intermediate", 5}
			if DockerDeepDive != want {
				t.Errorf("printVar want/got \n%v \n%v", want, DockerDeepDive)
			}
		})
}
