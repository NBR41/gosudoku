package model

import (
	"testing"
)

var inputs = []Input{
	{Coord: NewCoord(0, 0), V: 9},
	{Coord: NewCoord(3, 0), V: 1},
	{Coord: NewCoord(8, 0), V: 5},
	{Coord: NewCoord(2, 1), V: 5},
	{Coord: NewCoord(4, 1), V: 9},
	{Coord: NewCoord(6, 1), V: 2},
	{Coord: NewCoord(8, 1), V: 1},
	{Coord: NewCoord(0, 2), V: 8},
	{Coord: NewCoord(4, 2), V: 4},
	{Coord: NewCoord(4, 3), V: 8},
	{Coord: NewCoord(3, 4), V: 7},
	{Coord: NewCoord(4, 5), V: 2},
	{Coord: NewCoord(5, 5), V: 6},
	{Coord: NewCoord(8, 5), V: 9},
	{Coord: NewCoord(0, 6), V: 2},
	{Coord: NewCoord(3, 6), V: 3},
	{Coord: NewCoord(8, 6), V: 6},
	{Coord: NewCoord(3, 7), V: 2},
	{Coord: NewCoord(6, 7), V: 9},
	{Coord: NewCoord(2, 8), V: 1},
	{Coord: NewCoord(3, 8), V: 9},
	{Coord: NewCoord(5, 8), V: 4},
	{Coord: NewCoord(6, 8), V: 5},
	{Coord: NewCoord(7, 8), V: 7},
}

type CellTest struct {
	v   *int
	key string
}

func (c *CellTest) Get() *int {
	return c.v
}
func (c *CellTest) Set(v *int) {
	c.v = v
}
func (c *CellTest) Key() string {
	return c.key
}

func NewCellTest(x, y int) Celler {
	return &CellTest{key: Key(x, y)}
}

func TestDisplay(t *testing.T) {
	g, _ := NewGrid(9, NewCellTest)
	exp := `=========================================
||   |   |   ||   |   |   ||   |   |   ||
-----------------------------------------
||   |   |   ||   |   |   ||   |   |   ||
-----------------------------------------
||   |   |   ||   |   |   ||   |   |   ||
=========================================
||   |   |   ||   |   |   ||   |   |   ||
-----------------------------------------
||   |   |   ||   |   |   ||   |   |   ||
-----------------------------------------
||   |   |   ||   |   |   ||   |   |   ||
=========================================
||   |   |   ||   |   |   ||   |   |   ||
-----------------------------------------
||   |   |   ||   |   |   ||   |   |   ||
-----------------------------------------
||   |   |   ||   |   |   ||   |   |   ||
=========================================
`
	v := g.Display()
	if exp != v {
		t.Errorf("unexpected display")
	}

	g.Fill(inputs)
	exp = `=========================================
|| 9 |   |   || 1 |   |   ||   |   | 5 ||
-----------------------------------------
||   |   | 5 ||   | 9 |   || 2 |   | 1 ||
-----------------------------------------
|| 8 |   |   ||   | 4 |   ||   |   |   ||
=========================================
||   |   |   ||   | 8 |   ||   |   |   ||
-----------------------------------------
||   |   |   || 7 |   |   ||   |   |   ||
-----------------------------------------
||   |   |   ||   | 2 | 6 ||   |   | 9 ||
=========================================
|| 2 |   |   || 3 |   |   ||   |   | 6 ||
-----------------------------------------
||   |   |   || 2 |   |   || 9 |   |   ||
-----------------------------------------
||   |   | 1 || 9 |   | 4 || 5 | 7 |   ||
=========================================
`
	v = g.Display()
	if exp != v {
		t.Errorf("unexpected display")
	}
}

func TestNewGrid(t *testing.T) {
	if _, err := NewGrid(5, nil); err != ErrInvalidSquare {
		t.Errorf("expecting error: %v", ErrInvalidSquare)
	}

	g, err := NewGrid(9, NewCellTest)
	if err != nil {
		t.Errorf("unexpected error: [%v]", err)
	} else {
		if g.GetMax() != 9 {
			t.Errorf("unexpected GetMax value, exp [%d] got [%d]", 9, g.GetMax())
		}
		if g.GetBlockWidth() != 3 {
			t.Errorf("unexpected GetMax value, exp [%d] got [%d]", 3, g.GetBlockWidth())
		}
	}
	if _, err := NewGrid(5, nil); err != ErrInvalidSquare {
		t.Errorf("expecting error: %v", ErrInvalidSquare)
	}
}
