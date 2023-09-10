package maze

import (
  "testing"
)

func TestContains(t *testing.T) {
  arr := []int{1, 3, 5, 6} 
  elem := 5
  res := Contains(arr, elem)

  if !res {
    t.Fatalf("Array %v does not contain %d", arr, elem)
  }
}

func TestDoesntContain(t *testing.T) {
  arr := []int{1, 2, 3}
  elem := 4
  res := Contains(arr, elem)

  if res {
    t.Fatalf("Array %v contains %d", arr, elem)
  }
}
