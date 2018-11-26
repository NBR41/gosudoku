package model

import (
	"testing"
)

var inputs = []Input{
	{Coord: Coord{X: 0, Y: 0}, V: 9},
	{Coord: Coord{X: 3, Y: 0}, V: 1},
	{Coord: Coord{X: 8, Y: 0}, V: 5},
	{Coord: Coord{X: 2, Y: 1}, V: 5},
	{Coord: Coord{X: 4, Y: 1}, V: 9},
	{Coord: Coord{X: 6, Y: 1}, V: 2},
	{Coord: Coord{X: 8, Y: 1}, V: 1},
	{Coord: Coord{X: 0, Y: 2}, V: 8},
	{Coord: Coord{X: 4, Y: 2}, V: 4},
	{Coord: Coord{X: 4, Y: 3}, V: 8},
	{Coord: Coord{X: 3, Y: 4}, V: 7},
	{Coord: Coord{X: 4, Y: 5}, V: 2},
	{Coord: Coord{X: 5, Y: 5}, V: 6},
	{Coord: Coord{X: 8, Y: 5}, V: 9},
	{Coord: Coord{X: 0, Y: 6}, V: 2},
	{Coord: Coord{X: 3, Y: 6}, V: 3},
	{Coord: Coord{X: 8, Y: 6}, V: 6},
	{Coord: Coord{X: 3, Y: 7}, V: 2},
	{Coord: Coord{X: 6, Y: 7}, V: 9},
	{Coord: Coord{X: 2, Y: 8}, V: 1},
	{Coord: Coord{X: 3, Y: 8}, V: 9},
	{Coord: Coord{X: 5, Y: 8}, V: 4},
	{Coord: Coord{X: 6, Y: 8}, V: 5},
	{Coord: Coord{X: 7, Y: 8}, V: 7},
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
