package maze

import (
  "testing"
)

func TestUnvisitedNeighbours(t *testing.T) {
  dfs := DfsMazeGenerator {
    grid: Grid { height: 5, width: 5 },
    visited: []Point{ {0,1}, {1,2}},
  }

  p := Point { 1, 1}
  unvisitedNeighbours := dfs.unvisitedNeighbours(p)

  numberOfUnvisitedNeighbours := len(unvisitedNeighbours)
  if numberOfUnvisitedNeighbours != 2 {
    t.Errorf("Point %v should have 2 unvisited neighbours but has %d: %v", p, numberOfUnvisitedNeighbours, unvisitedNeighbours)
  }

  
}
