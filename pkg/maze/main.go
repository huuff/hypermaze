package maze

func main() {
  grid := Grid { Width: 5, Height: 5 }
  dfs := NewDfsMazeGenerator(grid)
  dfs.GenerateMaze()
} 
