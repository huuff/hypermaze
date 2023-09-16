package maze

import (
	"fmt"

	"github.com/samber/lo"
)

func (m Maze) AsciiView() {
  for y := range lo.Range((m.Grid.Height * 2) + 1) {
    for x := range lo.Range((m.Grid.Width * 2) + 1) {
      if x % 2 == 0 || y % 2 == 0 {
        fmt.Print("#")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n")
  }
}
