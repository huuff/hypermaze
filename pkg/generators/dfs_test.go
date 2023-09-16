package generators

import (
  "testing"
  "reflect"
  "xyz.haff/maze/pkg/grid"
)

func TestUnvisitedNeighbours(t *testing.T) {
  dfs := DfsMazeGenerator {
    grid: grid.Grid { Height: 5, Width: 5 },
    visited: []grid.Point{ {0,1}, {1,2}},
  }

  p := grid.Point { 1, 1}
  unvisitedNeighbours := dfs.unvisitedNeighbours(p)

  expected := []grid.Point{ {2, 1}, { 1, 0}}

  if !reflect.DeepEqual(expected, unvisitedNeighbours) {
    t.Fatalf("Expected %v, got %v", expected, unvisitedNeighbours)
  }

  
}
