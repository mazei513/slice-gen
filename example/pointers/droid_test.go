package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}
	expectedA := getExpectedA(a)
	//WHEN
	newA := DeleteDroid(1, a)
	//THEN
	expectedNewA := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestDeleteRangeDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
		&Droid{Name: "Do-Bot"},
	}
	expectedA := getExpectedA(a)
	//WHEN
	newA := DeleteRangeDroid(1, 2, a)
	//THEN
	expectedNewA := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Do-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestInsertDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}
	expectedA := getExpectedA(a)
	o := &Droid{Name: "Do-Bot"}
	//WHEN
	newA := InsertDroid(2, o, a)
	//THEN
	expectedNewA := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Do-Bot"},
		&Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestPushDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}
	expectedA := getExpectedA(a)
	o := &Droid{Name: "Do-Bot"}
	//WHEN
	newA := PushDroid(o, a)
	//THEN
	expectedNewA := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
		&Droid{Name: "Do-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestPopDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}
	expectedA := getExpectedA(a)
	//WHEN
	o, newA := PopDroid(a)
	//THEN
	expectedO := &Droid{Name: "Fo-Bot"}
	expectedNewA := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedO, o)
	assert.Equal(t, expectedNewA, newA)
}

func TestUnshiftDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}
	expectedA := getExpectedA(a)
	o := &Droid{Name: "Do-Bot"}
	//WHEN
	newA := UnshiftDroid(o, a)
	//THEN
	expectedNewA := []*Droid{
		&Droid{Name: "Do-Bot"},
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestShiftDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot"},
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}
	expectedA := getExpectedA(a)
	//WHEN
	o, newA := ShiftDroid(a)
	//THEN
	expectedO := &Droid{Name: "Bo-Bot"}
	expectedNewA := []*Droid{
		&Droid{Name: "Go-Bot"},
		&Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedO, o)
	assert.Equal(t, expectedNewA, newA)
}

func TestFilterDroid(t *testing.T) {
	//GIVEN
	a := []*Droid{
		&Droid{Name: "Bo-Bot", Version: 1},
		&Droid{Name: "Go-Bot", Version: 2},
		&Droid{Name: "Fo-Bot", Version: 3},
	}
	expectedA := getExpectedA(a)
	fn := func(d *Droid) bool {
		return d.Version > 1
	}
	//WHEN
	newA := FilterDroid(a, fn)
	//THEN
	expectedNewA := []*Droid{
		&Droid{Name: "Go-Bot", Version: 2},
		&Droid{Name: "Fo-Bot", Version: 3},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func getExpectedA(a []*Droid) []*Droid {
	expectedA := make([]*Droid, len(a))
	copy(expectedA, a)
	return expectedA
}
