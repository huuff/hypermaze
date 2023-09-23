package main

import (
	"fmt"

	"xyz.haff/maze/pkg/direction"
)

func directionToString(d direction.Direction) string {
  switch (d) {
    case direction.North:
      return "North"
    case direction.South:
      return "South"
    case direction.East:
      return "East"
    case direction.West:
      return "West"
  }
  return fmt.Sprintf("Weird direction %d", d)
}
