package backtracking

import (
	"reflect"
	"testing"

	"github.com/NBR41/gosudoku/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestNewNode(t *testing.T) {
	v := 10
	n := &node{Coord: model.NewCoord(0, 0), k: "0-0"}
	if n.Get() != nil {
		t.Error("unexpected value")
	}
	n.Set(&v)
	if n.Get() != &v {
		t.Error("unexpected value")
	}
	if n.Key() != "0-0" {
		t.Errorf("unexpected value, exp [%s] got [%s]", "0-0", n.Key())
	}
}

func TestIsValid(t *testing.T) {
	v := 10
	n := &node{Coord: model.NewCoord(0, 0), k: "0-0"}
	if n.IsValid() {
		t.Error("unexpected value")
	}
	n.v = &v
	if !n.IsValid() {
		t.Error("unexpected value")
	}
	n.nodes = [][]model.Celler{
		{
			&node{
				Coord: model.NewCoord(0, 0),
				k:     "0-0",
				v:     &v,
			},
		},
	}
	if n.IsValid() {
		t.Error("unexpected value")
	}
}

func TestSetNodes(t *testing.T) {
	nodes := make(map[string]model.Celler)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			nodes[model.Key(i, j)] = &node{Coord: model.NewCoord(i, j), k: model.Key(i, j)}
		}
	}
	n := &node{Coord: model.NewCoord(0, 0), k: "0-0"}
	n.setNodes(nodes, 4, 2)
	exp := [][]model.Celler{
		{nodes["1-0"], nodes["2-0"], nodes["3-0"]},
		{nodes["0-1"], nodes["0-2"], nodes["0-3"]},
		{nodes["0-1"], nodes["1-0"], nodes["1-1"]},
	}
	if diff := pretty.Compare(exp, n.nodes); diff != "" {
		t.Errorf("diff (-got +want)\n%s", diff)
	}
}

func TestNewGrid(t *testing.T) {
	g, err := NewGrid(3)
	if err == nil {
		t.Error("expecting error")
	} else if err != model.ErrInvalidSquare {
		t.Errorf("unexpected error, exp [%v] got [%v]", model.ErrInvalidSquare, err)
	}
	if g != nil {
		t.Error("unexpected value")
	}

	g, _ = NewGrid(4)
	exp, _ := model.NewGrid(4, newNode)
	for k := range exp.Values {
		exp.Values[k].(*node).setNodes(exp.Values, 4, 2)
	}
	if !reflect.DeepEqual(exp, g) {
		t.Error("unexpected value")
	}
}
