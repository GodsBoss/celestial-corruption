package game

import (
  "math"
)

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

func distance(x1, y1 float64, x2, y2 float64) float64 {
  return math.Sqrt((x1-x2)*(x1-x2)+(y1-y2)*(y1-y2))
}
