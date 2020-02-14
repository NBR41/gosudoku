# gosudoku
![Code validation](https://github.com/NBR41/gosudoku/workflows/Code%20validation/badge.svg)
[![codecov](https://codecov.io/gh/NBR41/gosudoku/branch/master/graph/badge.svg)](https://codecov.io/gh/NBR41/gosudoku)
![Publish](https://github.com/NBR41/gosudoku/workflows/Publish/badge.svg)

A sudoku resolver

## Usage
```
Usage:
   [flags]

Flags:
  -h, --help          help for this command
      --path string   path of the file with coordinates
      --size int      grid number of cells [4,9,16,25] (default 9)
      --verbose       verbose resolution
      --yaml          is the input file a YAML
```

## Docker

```
docker run --rm -v ~/go/src/github.com/NBR41/gosudoku/testdata/grid.json:/tmp/grid.json nbr41/gosudoku:latest --path=/tmp/grid.json
```
