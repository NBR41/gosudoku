package backtracking

import (
	"testing"

	"github.com/NBR41/gosudoku/model"
)

func TestResolve(t *testing.T) {
	var inputs = []model.Input{
		{Coord: model.NewCoord(0, 0), V: 9},
		{Coord: model.NewCoord(3, 0), V: 1},
		{Coord: model.NewCoord(8, 0), V: 5},
		{Coord: model.NewCoord(2, 1), V: 5},
		{Coord: model.NewCoord(4, 1), V: 9},
		{Coord: model.NewCoord(6, 1), V: 2},
		{Coord: model.NewCoord(8, 1), V: 1},
		{Coord: model.NewCoord(0, 2), V: 8},
		{Coord: model.NewCoord(4, 2), V: 4},
		{Coord: model.NewCoord(4, 3), V: 8},
		{Coord: model.NewCoord(3, 4), V: 7},
		{Coord: model.NewCoord(4, 5), V: 2},
		{Coord: model.NewCoord(5, 5), V: 6},
		{Coord: model.NewCoord(8, 5), V: 9},
		{Coord: model.NewCoord(0, 6), V: 2},
		{Coord: model.NewCoord(3, 6), V: 3},
		{Coord: model.NewCoord(8, 6), V: 6},
		{Coord: model.NewCoord(3, 7), V: 2},
		{Coord: model.NewCoord(6, 7), V: 9},
		{Coord: model.NewCoord(2, 8), V: 1},
		{Coord: model.NewCoord(3, 8), V: 9},
		{Coord: model.NewCoord(5, 8), V: 4},
		{Coord: model.NewCoord(6, 8), V: 5},
		{Coord: model.NewCoord(7, 8), V: 7},
	}

	g, _ := NewGrid(9)
	g.Fill(inputs)
	var subs int
	err := Resolve(g, PostProcess(func(g *model.Grid) {
		subs++
	}))
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	}

	for i := range g.Values {
		if g.Values[i].Get() == nil {
			t.Error("unexpected value")
			break
		}

		if !g.Values[i].(*node).IsValid() {
			t.Error("unexpected value")
			break
		}
	}

	if subs == 0 {
		t.Error("unexpected value")
	}
}
