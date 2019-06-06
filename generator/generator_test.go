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

func TestGetPackage(t *testing.T) {
	testCases := []struct {
		desc     string
		dir      string
		hasError bool
		expected string
	}{
		{
			desc:     "normal",
			dir:      ".",
			expected: "generator",
		},
		{
			desc:     "non-existent directory",
			dir:      "./foo",
			hasError: true,
		},
		{
			desc:     "dir with no packages",
			dir:      "./../..",
			hasError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			pkg, err := getPackage(tC.dir)
			if tC.hasError {
				assert.Error(t, err)
				t.Logf("%+v", err)
				assert.Nil(t, pkg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.expected, pkg.Name)
			}
		})
	}
}

func TestGetFilename(t *testing.T) {
	testCases := []struct {
		desc     string
		dir      string
		filename string
		expected string
	}{
		{
			desc:     "absolute",
			dir:      "/path/to/dir",
			filename: "robot",
			expected: "/path/to/dir/robot_slicegen.go",
		},
		{
			desc:     "relative",
			dir:      "dir",
			filename: "robot",
			expected: "dir/robot_slicegen.go",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			fp := getFilepath(tC.dir, tC.filename)
			assert.Equal(t, tC.expected, fp)
		})
	}
}
