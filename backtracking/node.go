package backtracking

import (
	"github.com/NBR41/gosudoku/model"
)

//NewGrid returns new instance of grid
func NewGrid(n int) (*model.Grid, error) {
	g, err := model.NewGrid(n, newNode)
	if err != nil {
		return nil, err
	}

	for k := range g.Values {
		g.Values[k].(*node).setNodes(g.Values, g.GetMax(), g.GetBlockWidth())
	}
	return g, nil
}

//node a box of the grid
type node struct {
	model.Coord
	k string
	v *int

	nodes [][]model.Celler
}

func newNode(x, y int) model.Celler {
	return &node{Coord: model.NewCoord(x, y), k: model.Key(x, y)}
}

func (n *node) Get() *int {
	return n.v
}

func (n *node) Set(v *int) {
	n.v = v
}

func (n *node) Key() string {
	return n.k
}

func (n *node) setNodes(nodes map[string]model.Celler, max, blocNb int) {
	sliceLen := max - 1
	xs := make([]model.Celler, sliceLen)
	nb := 0
	for x := 0; x < max; x++ {
		if x != n.X() {
			xs[nb] = nodes[model.Key(x, n.Y())]
			nb++
		}
	}
	ys := make([]model.Celler, sliceLen)
	nb = 0
	for y := 0; y < max; y++ {
		if y != n.Y() {
			ys[nb] = nodes[model.Key(n.X(), y)]
			nb++
		}
	}

	bs := make([]model.Celler, sliceLen)
	ox := n.X() - (n.X() % blocNb)
	oy := n.Y() - (n.Y() % blocNb)
	i := 0
	for x := ox; x < ox+blocNb; x++ {
		for y := oy; y < oy+blocNb; y++ {
			if n.X() == x && n.Y() == y {
				continue
			}
			bs[i] = nodes[model.Key(x, y)]
			i++
		}
	}

	n.nodes = [][]model.Celler{xs, ys, bs}
}

//IsValid check if the value is valid in the grid
func (n *node) IsValid() bool {
	if n.v == nil {
		return false
	}
	for i := range n.nodes {
		for j := range n.nodes[i] {
			v := n.nodes[i][j].Get()
			if v != nil && *v == *n.v {
				return false
			}
		}
	}
	return true
}
