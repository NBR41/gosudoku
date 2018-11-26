package backtracking

import (
	"testing"

	"github.com/NBR41/gosudoku/model"
)

var inputs = []model.Input{
	{Coord: model.Coord{X: 0, Y: 0}, V: 9},
	{Coord: model.Coord{X: 3, Y: 0}, V: 1},
	{Coord: model.Coord{X: 8, Y: 0}, V: 5},
	{Coord: model.Coord{X: 2, Y: 1}, V: 5},
	{Coord: model.Coord{X: 4, Y: 1}, V: 9},
	{Coord: model.Coord{X: 6, Y: 1}, V: 2},
	{Coord: model.Coord{X: 8, Y: 1}, V: 1},
	{Coord: model.Coord{X: 0, Y: 2}, V: 8},
	{Coord: model.Coord{X: 4, Y: 2}, V: 4},
	{Coord: model.Coord{X: 4, Y: 3}, V: 8},
	{Coord: model.Coord{X: 3, Y: 4}, V: 7},
	{Coord: model.Coord{X: 4, Y: 5}, V: 2},
	{Coord: model.Coord{X: 5, Y: 5}, V: 6},
	{Coord: model.Coord{X: 8, Y: 5}, V: 9},
	{Coord: model.Coord{X: 0, Y: 6}, V: 2},
	{Coord: model.Coord{X: 3, Y: 6}, V: 3},
	{Coord: model.Coord{X: 8, Y: 6}, V: 6},
	{Coord: model.Coord{X: 3, Y: 7}, V: 2},
	{Coord: model.Coord{X: 6, Y: 7}, V: 9},
	{Coord: model.Coord{X: 2, Y: 8}, V: 1},
	{Coord: model.Coord{X: 3, Y: 8}, V: 9},
	{Coord: model.Coord{X: 5, Y: 8}, V: 4},
	{Coord: model.Coord{X: 6, Y: 8}, V: 5},
	{Coord: model.Coord{X: 7, Y: 8}, V: 7},
}

func TestGetPossibilities(t *testing.T) {
	g, _ := NewGrid(4)
	//g.Fill(inputs)
	pos := getPossibilities(g.Values, g.Max)
	//exp := []*Possibilities{}
	t.Log(len(pos))
	for i := range pos {
		t.Log(pos[i])
	}
}

func TestResolve(t *testing.T) {
	g, _ := NewGrid(9)
	g.Fill(inputs)
	Resolve(g)
}
