//go:generate go run github.com/mazei513/slice-gen testObj

package tests

type testObj struct {
	f1 string
	f2 int
	f3 *string
}
