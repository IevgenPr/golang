package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ipr0/golang/models"
)

func TestAddUser(t *testing.T) {
	// u := &models.User{FirstName: "First", LastName: "Last"}
	// msg := "New user must not include IDor it must be set to zero."
	cases := []struct {
		in, want models.User
		err      error
	}{
		// case 1
		{models.User{FirstName: "First", LastName: "Last"},
			models.User{ID: 1, FirstName: "First", LastName: "Last"}, nil},
		// case 2 - exception
		{models.User{
			ID:        1,
			FirstName: "Second",
			LastName:  "Last",
		},
			models.User{}, models.ErrBadUser},
		//case 3 - empty
		{models.User{}, models.User{}, models.ErrBadUser},
	}
	for _, c := range cases {
		got, err := models.AddUser(c.in)
		// using cmp.Equal as more accurate to reflect one
		if !cmp.Equal(got, c.want) {
			t.Errorf("\nAddUser(%q) == %q, want %q", c.in, got, c.want)
		}

		if err != c.err {
			t.Errorf("\nAddUser(%q) == err: %q, want: %q", c.in, err, c.err)
		}
	}
}
