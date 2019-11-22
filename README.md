# gosudoku
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
