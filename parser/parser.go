// Package parser provides functions to create a slice of inputs from JSon or YAML
package parser

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/NBR41/gosudoku/model"

	yaml "gopkg.in/yaml.v2"
)

type decoder interface {
	Decode(interface{}) error
}

type cell struct {
	X int `json:"x" yaml:"x"`
	Y int `json:"y" yaml:"y"`
	V int `json:"v" yaml:"v"`
}

// NewGridFromJSON returns a slice of inputs for a JSON reader
func NewGridFromJSON(r io.Reader) ([]model.Input, error) {
	return parse(json.NewDecoder(r))
}

// NewGridFromYAML returns a slice of inputs for a YAML reader
func NewGridFromYAML(r io.Reader) ([]model.Input, error) {
	return parse(yaml.NewDecoder(r))
}

func parse(dec decoder) ([]model.Input, error) {
	var (
		c   []cell
		err error
	)
	for {
		err = dec.Decode(&c)
		switch {
		case err != nil:
			return nil, fmt.Errorf("unable to parse cells: %w", err)
		default:
			ret := make([]model.Input, len(c))
			for i := range c {
				ret[i] = model.Input{Coord: model.NewCoord(c[i].X, c[i].Y), V: c[i].V}
			}
			return ret, nil
		}
	}
}
