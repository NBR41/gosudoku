package parser

import (
	"errors"
	"strings"
	"testing"

	"github.com/NBR41/gosudoku/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	exp := []model.Input{
		{Coord: model.NewCoord(0, 0), V: 9},
		{Coord: model.NewCoord(3, 0), V: 1},
		{Coord: model.NewCoord(8, 0), V: 5},
	}

	t.Run("decode error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := NewMockdecoder(ctrl)
		mock.EXPECT().Decode(gomock.Any()).Return(errors.New("decode error"))
		v, err := parse(mock)
		assert.EqualError(t, err, "unable to parse cells: decode error")
		assert.Nil(t, v)
	})

	t.Run("json", func(t *testing.T) {
		var jsonStream = `[
{"x":0,"y":0,"v":9},
{"x":3,"y":0,"v":1},
{"x":8,"y":0,"v":5}
]`
		v, err := NewGridFromJSON(strings.NewReader(jsonStream))
		assert.Nil(t, err)

		assert.Equal(t, exp, v)
	})

	t.Run("yaml", func(t *testing.T) {
		var jsonStream = `- x: 0
  y: 0
  v: 9
- x: 3
  y: 0
  v: 1
- x: 8
  y: 0
  v: 5`
		v, err := NewGridFromYAML(strings.NewReader(jsonStream))
		assert.Nil(t, err)
		assert.Equal(t, exp, v)
	})
}
