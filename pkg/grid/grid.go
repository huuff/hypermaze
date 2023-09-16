package grid

type Point struct {
  X int
  Y int
}

type Grid struct {
  Height int
  Width int
}

func (g Grid) Neighbours(p Point) []Point {
  neighbours := make([]Point, 0)
  
  // Clockwork from top
  if p.Y < (g.Height - 1) {
    neighbours = append(neighbours, Point { p.X, p.Y + 1 })
  }
  
  if p.X < (g.Width - 1) {
    neighbours = append(neighbours, Point { p.X + 1, p.Y })
  }

  if p.Y > 0 {
    neighbours = append(neighbours, Point { p.X, p.Y - 1})
  }

  if p.X > 0 {
    neighbours = append(neighbours, Point { p.X - 1, p.Y })
  }

  return neighbours
}
