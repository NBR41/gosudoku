package backtracking

import (
	"fmt"
	"time"

	"github.com/NBR41/gosudoku/model"
)

//PostProcess type for post process function
type PostProcess func(g *model.Grid)

//WithDelayedDisplay display the grid and wait d to continue
func WithDelayedDisplay(d time.Duration) PostProcess {
	return func(g *model.Grid) {
		fmt.Println(g.Display())
		<-time.After(d)
	}
}

//Resolve resolve the sudoku
func Resolve(g *model.Grid, fs ...PostProcess) error {
	//var nodes map[string]model.Celler
	if processPos(g, getPossibilities(g.Values, g.GetMax()), 0, fs) == false {
		return model.ErrNoSolution
	}
	return nil
}

func processPos(g *model.Grid, pos []*possibilities, n int, fs []PostProcess) bool {
	if n == len(pos) {
		return true
	}
	for i := range pos[n].Values {
		pos[n].Node.v = &pos[n].Values[i]

		if pos[n].Node.IsValid() == true {
			for j := range fs {
				fs[j](g)
			}

			if processPos(g, pos, n+1, fs) {
				return true
			}
		}
	}
	pos[n].Node.v = nil
	return false
}
