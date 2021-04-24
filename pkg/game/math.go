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

func normalizedSpeed(speed float64, ms int) float64 {
  return speed * float64(ms) / 1000.0
}
