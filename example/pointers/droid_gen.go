// Code generated by github.com/mazei513/slice-gen, DO NOT EDIT.
package pointers

func DeleteDroid(i int, a []*Droid) []*Droid {
	s := make([]*Droid, i, len(a)-1)
	copy(s, a)
	return append(s, a[i+1:]...)
}
func DeleteRangeDroid(from, to int, a []*Droid) []*Droid {
	s := make([]*Droid, from, len(a)-(to-from))
	copy(s, a)
	return append(s, a[to+1:]...)
}
func InsertDroid(i int, o *Droid, a []*Droid) []*Droid {
	var zero *Droid
	s := append(a, zero)
	copy(s[i+1:], a[i:])
	s[i] = o
	return s
}
func PushDroid(o *Droid, a []*Droid) []*Droid {
	return append(a, o)
}
func PopDroid(a []*Droid) (*Droid, []*Droid) {
	return a[len(a)-1], append([]*Droid(nil), a[:len(a)-1]...)
}
func UnshiftDroid(o *Droid, a []*Droid) []*Droid {
	return append([]*Droid{o}, a...)
}
func ShiftDroid(a []*Droid) (*Droid, []*Droid) {
	return a[0], append([]*Droid(nil), a[1:]...)
}
func FilterDroid(a []*Droid, f func(*Droid) bool) []*Droid {
	b := a[:0]
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}
