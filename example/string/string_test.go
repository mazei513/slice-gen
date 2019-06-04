package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	//WHEN
	newA := DeleteString(1, a)
	//THEN
	expectedA := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	expectedNewA := []string{
		"Boo",
		"Foo",
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestDeleteRangeString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Goo",
		"Foo",
		"Doo",
	}
	//WHEN
	newA := DeleteRangeString(1, 2, a)
	//THEN
	expectedA := []string{
		"Boo",
		"Goo",
		"Foo",
		"Doo",
	}
	expectedNewA := []string{
		"Boo",
		"Doo",
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestInsertString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	o := "Doo"
	//WHEN
	newA := InsertString(2, o, a)
	//THEN
	expectedA := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	expectedNewA := []string{
		"Boo",
		"Goo",
		"Doo",
		"Foo",
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestPushString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	o := "Doo"
	//WHEN
	newA := PushString(o, a)
	//THEN
	expectedA := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	expectedNewA := []string{
		"Boo",
		"Goo",
		"Foo",
		"Doo",
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestPopString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	//WHEN
	o, newA := PopString(a)
	//THEN
	expectedA := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	expectedO := "Foo"
	expectedNewA := []string{
		"Boo",
		"Goo",
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedO, o)
	assert.Equal(t, expectedNewA, newA)
}

func TestUnshiftString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	o := "Doo"
	//WHEN
	newA := UnshiftString(o, a)
	//THEN
	expectedA := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	expectedNewA := []string{
		"Doo",
		"Boo",
		"Goo",
		"Foo",
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestShiftString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	//WHEN
	o, newA := ShiftString(a)
	//THEN
	expectedA := []string{
		"Boo",
		"Goo",
		"Foo",
	}
	expectedO := "Boo"
	expectedNewA := []string{
		"Goo",
		"Foo",
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedO, o)
	assert.Equal(t, expectedNewA, newA)
}

func TestFilterString(t *testing.T) {
	//GIVEN
	a := []string{
		"Boo",
		"Gooo",
		"Foooo",
	}
	fn := func(s string) bool {
		return len(s) > 3
	}
	//WHEN
	newA := FilterString(a, fn)
	//THEN
	expectedA := []string{
		"Gooo",
		"Foooo",
	}

	assert.Equal(t, expectedA, newA)
}
