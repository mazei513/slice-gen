package droid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	//WHEN
	newA := DeleteDroid(1, a)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	expectedNewA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestDeleteRangeDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
		Droid{Name: "Do-Bot"},
	}
	//WHEN
	newA := DeleteRangeDroid(1, 2, a)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
		Droid{Name: "Do-Bot"},
	}
	expectedNewA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Do-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestInsertDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	o := Droid{Name: "Do-Bot"}
	//WHEN
	newA := InsertDroid(2, o, a)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	expectedNewA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Do-Bot"},
		Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestPushDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	o := Droid{Name: "Do-Bot"}
	//WHEN
	newA := PushDroid(o, a)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	expectedNewA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
		Droid{Name: "Do-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestPopDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	//WHEN
	o, newA := PopDroid(a)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	expectedO := Droid{Name: "Fo-Bot"}
	expectedNewA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedO, o)
	assert.Equal(t, expectedNewA, newA)
}

func TestUnshiftDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	o := Droid{Name: "Do-Bot"}
	//WHEN
	newA := UnshiftDroid(o, a)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	expectedNewA := []Droid{
		Droid{Name: "Do-Bot"},
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedNewA, newA)
}

func TestShiftDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	//WHEN
	o, newA := ShiftDroid(a)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Bo-Bot"},
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}
	expectedO := Droid{Name: "Bo-Bot"}
	expectedNewA := []Droid{
		Droid{Name: "Go-Bot"},
		Droid{Name: "Fo-Bot"},
	}

	assert.Equal(t, expectedA, a)
	assert.Equal(t, expectedO, o)
	assert.Equal(t, expectedNewA, newA)
}

func TestFilterDroid(t *testing.T) {
	//GIVEN
	a := []Droid{
		Droid{Name: "Bo-Bot", Version: 1},
		Droid{Name: "Go-Bot", Version: 2},
		Droid{Name: "Fo-Bot", Version: 3},
	}
	fn := func(d Droid) bool {
		return d.Version > 1
	}
	//WHEN
	newA := FilterDroid(a, fn)
	//THEN
	expectedA := []Droid{
		Droid{Name: "Go-Bot", Version: 2},
		Droid{Name: "Fo-Bot", Version: 3},
	}

	assert.Equal(t, expectedA, newA)
}
