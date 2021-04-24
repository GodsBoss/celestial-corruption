package game

func max(value int, rest ...int) int {
  for i := range rest {
    if rest[i] > value {
      value = rest[i]
    }
  }
  return value
}

func min(value int, rest ...int) int {
  for i := range rest {
    if rest[i] < value {
      value = rest[i]
    }
  }
  return value
}
