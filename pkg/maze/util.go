package maze

func Contains[T comparable](arr []T, elem T) bool {
  for _, curr := range arr {
    if curr == elem {
      return true
    }
  }

  return false
}
