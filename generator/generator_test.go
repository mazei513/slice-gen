package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNameFromObjType(t *testing.T) {
	cases := []struct {
		desc     string
		s        string
		expected string
	}{
		{"exported", "Droid", "Droid"},
		{"unexported", "droid", "Droid"},
		{"exported pointer", "*Droid", "Droid"},
		{"unexported pointer", "*droid", "Droid"},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			a := getNameFromObjType(c.s)
			assert.Equal(t, c.expected, a)
		})
	}
}
