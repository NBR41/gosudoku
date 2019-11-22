package backtracking

import (
	"reflect"
	"testing"

	"github.com/NBR41/gosudoku/model"
)

func TestGetMapValues(t *testing.T) {
	v := getMapValues(4)
	exp := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}
	if !reflect.DeepEqual(exp, v) {
		t.Error("unexpected value")
	}
}

func TestGetPossibilities(t *testing.T) {
	cells := map[string]model.Celler{
		"0-0": newNode(0, 0),
		"0-1": newNode(0, 1),
		"1-0": newNode(1, 0),
		"1-1": newNode(1, 1),
	}
	v := 1
	cells["0-0"].Set(&v)
	cells["0-0"].(*node).nodes = [][]model.Celler{{cells["0-1"]}, {cells["1-0"]}}
	cells["0-1"].(*node).nodes = [][]model.Celler{{cells["0-0"]}, {cells["1-1"]}}
	cells["1-0"].(*node).nodes = [][]model.Celler{{cells["0-0"]}, {cells["1-1"]}}
	cells["1-1"].(*node).nodes = [][]model.Celler{{cells["0-1"]}, {cells["1-0"]}}

	pos := getPossibilities(cells, 2)
	if len(pos) != 3 {
		t.Error("unexpected value")
	} else {
		if pos[0].Node.Key() != "0-1" && pos[0].Node.Key() != "1-0" && !reflect.DeepEqual(pos[0].Values, []int{2}) {
			t.Error("unexpected value")
		}
		if pos[1].Node.Key() != "0-1" && pos[1].Node.Key() != "1-0" && !reflect.DeepEqual(pos[1].Values, []int{2}) {
			t.Error("unexpected value")
		}
		if pos[2].Node.Key() != "1-1" && pos[2].Node.Key() != "1-0" && !reflect.DeepEqual(pos[2].Values, []int{1, 2}) {
			t.Error("unexpected value")
		}
	}
}
