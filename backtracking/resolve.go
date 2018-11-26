package backtracking

import (
	"fmt"
	"sort"

	"github.com/NBR41/gosudoku/model"
)

type possibilities struct {
	Node   *node
	Values []int
}

func getMapValues(max int) map[int]struct{} {
	ret := make(map[int]struct{})
	for i := 1; i <= max; i++ {
		ret[i] = struct{}{}
	}
	return ret
}

//Resolve resolve the sudoku
func Resolve(g *model.Grid) {
	//var nodes map[string]model.Celler
	if processPos(g, getPossibilities(g.Values, g.Max), 0) == false {
		fmt.Println("no solution")
	}
}

func getPossibilities(nodes map[string]model.Celler, max int) []*possibilities {
	var ret []*possibilities
	var no *node
	var v *int
	for k := range nodes {
		if nodes[k].Get() != nil {
			continue
		}

		vals := getMapValues(max)
		no = nodes[k].(*node)
		for i := range no.nodes {
			for j := range no.nodes[i] {
				v = no.nodes[i][j].Get()
				if v != nil {
					delete(vals, *v)
				}
			}
		}
		pos := &possibilities{Node: nodes[k].(*node)}
		for v := range vals {
			pos.Values = append(pos.Values, v)
		}
		ret = append(ret, pos)
	}
	sort.SliceStable(ret, func(i int, j int) bool {
		return len(ret[i].Values) < len(ret[j].Values)
	})
	return ret
}

func processPos(g *model.Grid, pos []*possibilities, n int) bool {
	if n == len(pos) {
		return true
	}
	for i := range pos[n].Values {
		pos[n].Node.v = &pos[n].Values[i]
		//fmt.Println(fmt.Sprintf("%+v", pos[n].Node))
		if pos[n].Node.IsValid() == true {
			fmt.Println(g.Display())
			//<-time.After(150 * time.Millisecond)
			if processPos(g, pos, n+1) {
				return true
			}
		}
	}
	pos[n].Node.v = nil
	return false
}
