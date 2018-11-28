package model

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Key return new key
func Key(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

//Coord coordonate of a cell in the grid
type Coord struct {
	x int
	y int
}

//X return x
func (c Coord) X() int {
	return c.x
}

//Y return Y
func (c Coord) Y() int {
	return c.y
}

//NewCoord return new Coord
func NewCoord(x, y int) Coord {
	return Coord{x: x, y: y}
}

//Input struct for an input in the cell
type Input struct {
	Coord
	V int
}

//ErrInvalidSquare error for invalid input
var ErrInvalidSquare = errors.New("invalid square")

//ErrNoSolution error for no solution
var ErrNoSolution = errors.New("no solution")

//Celler for a cell of the grid
type Celler interface {
	Get() *int
	Set(*int)
	Key() string
}

//Grid struct for sudoku grid
type Grid struct {
	Values     map[string]Celler
	max        int
	blockWidth int
}

//GetMax return the max number in the grid
func (g *Grid) GetMax() int {
	return g.max
}

//GetBlockWidth returns the block width of the grid
func (g *Grid) GetBlockWidth() int {
	return g.blockWidth
}

//NewGrid returns new instance of grid
func NewGrid(max int, f func(x, y int) Celler) (*Grid, error) {
	g := &Grid{
		Values:     make(map[string]Celler),
		max:        max,
		blockWidth: int(math.Sqrt(float64(max))),
	}

	if float64(g.blockWidth) != math.Sqrt(float64(max)) {
		return nil, ErrInvalidSquare
	}

	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			v := f(i, j)
			g.Values[v.Key()] = v
		}
	}
	return g, nil
}

//Display display the grid
func (g *Grid) Display() string {
	line := strings.Repeat("=", (4*g.max)+g.blockWidth+2) + "\n"
	subline := strings.Repeat("-", (4*g.max)+g.blockWidth+2) + "\n"
	output := line
	var v *int
	for i := 0; i < g.max; i++ {
		for j := 0; j < g.max; j++ {
			if j%g.blockWidth == 0 {
				output += "||"
			} else {
				output += "|"
			}
			v = g.Values[Key(j, i)].Get()
			if v == nil {
				output += "   "
			} else {
				output += " " + strconv.FormatInt(int64(*v), 10) + " "
			}
		}

		output += "||\n"
		if i%g.blockWidth == g.blockWidth-1 {
			output += line
		} else {
			output += subline
		}
	}
	return output
}

//Fill fill the grid with inputs
func (g *Grid) Fill(in []Input) {
	for i := range in {
		g.Values[Key(in[i].X(), in[i].Y())].Set(&in[i].V)
	}
}
