//go:generate go run github.com/mazei513/slice-gen Droid

package example

type Droid struct {
	Name    string
	Version int
	Owner   *string
}
