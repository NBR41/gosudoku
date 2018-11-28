package backtracking

import (
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
