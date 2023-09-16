package generators

import "xyz.haff/maze/pkg/grid"

type Passages = [][2]grid.Point

type PassageGenerator interface {
  GeneratePassages() Passages
}
