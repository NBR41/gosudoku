package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/NBR41/gosudoku/backtracking"
	"github.com/NBR41/gosudoku/model"
	"github.com/NBR41/gosudoku/parser"
	"github.com/spf13/cobra"
)

func main() {
	var (
		path            string
		size            int
		isYAML, verbose bool
	)

	cmd := &cobra.Command{
		Short: "Resolve a sudoku with coordinates configured by file",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if path == "" {
				return errors.New("empty file path")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := process(path, size, isYAML, verbose); err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.Flags().StringVar(&path, "path", "", "path of the file with coordinates")
	cmd.Flags().IntVar(&size, "size", 9, "grid number of cells [4,9,16,25]")
	cmd.Flags().BoolVar(&isYAML, "yaml", false, "is the input file a YAML")
	cmd.Flags().BoolVar(&verbose, "verbose", false, "verbose resolution")
	_ = cmd.MarkFlagRequired("path")

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func process(path string, size int, isYAML, verbose bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	var inputs []model.Input
	if isYAML {
		inputs, err = parser.NewGridFromYAML(f)
	} else {
		inputs, err = parser.NewGridFromJSON(f)
	}
	if err != nil {
		return err
	}

	g, err := backtracking.NewGrid(size)
	if err != nil {
		return err
	}
	g.Fill(inputs)

	fmt.Println(g.Display())

	var opts []backtracking.PostProcess
	if verbose {
		opts = append(opts, backtracking.WithDelayedDisplay(50*time.Millisecond))
	}
	if err = backtracking.Resolve(g, opts...); err != nil {
		return err
	}
	fmt.Println(g.Display())
	return nil
}
