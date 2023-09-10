package maze

import (
  "testing"
  "reflect"
)

func TestUnvisitedNeighbours(t *testing.T) {
  dfs := DfsMazeGenerator {
    grid: Grid { height: 5, width: 5 },
    visited: []Point{ {0,1}, {1,2}},
  }

  p := Point { 1, 1}
  unvisitedNeighbours := dfs.unvisitedNeighbours(p)

  expected := []Point{ {2, 1}, { 1, 0}}

  if !reflect.DeepEqual(expected, unvisitedNeighbours) {
    t.Fatalf("Expected %v, got %v", expected, unvisitedNeighbours)
  }

  
}
