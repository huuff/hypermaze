package maze

type Point struct {
  x int
  y int
}

type Grid struct {
  height int
  width int
}

func (g Grid) Neighbours(p Point) []Point {
  neighbours := make([]Point, 0)
  
  // Clockwork from top
  if p.y < (g.height - 1) {
    neighbours = append(neighbours, Point { p.x, p.y + 1 })
  }
  
  if p.x < (g.width - 1) {
    neighbours = append(neighbours, Point { p.x + 1, p.y })
  }

  if p.y > 0 {
    neighbours = append(neighbours, Point { p.x, p.y - 1})
  }

  if p.x > 0 {
    neighbours = append(neighbours, Point { p.x - 1, p.y })
  }

  return neighbours
}
