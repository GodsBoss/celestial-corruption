package game

import (
  "strings"
)

type text struct {
  contents string
  x int
  y int
}

func newText(contents string, x, y int) *text{
  return &text{
    contents: strings.Map(allowChar, strings.ToLower(contents)),
    x: x,
    y: y,
  }
}

func (txt *text) renderable(sf *spriteFactory) renderable {
  r := make(renderables, 0, len(txt.contents))
  lines := strings.Split(txt.contents, "\n")
  for i := range lines {
    for j := range lines[i] {
      r = append(
        r,
        sf.create("char_" + lines[i][j:j+1], txt.x + j * (charWidth+1), txt.y + i * (charHeight+1), 0),
      )
    }
  }
  return r
}

const allowedChars = "\nabcdefghijklmnopqrstuvwxyz 01234567890.,:!?()[]-=_\"'+"

var allowedCharsMap = func() map[rune]struct{} {
	m := map[rune]struct{}{}
	for _, r := range []rune(allowedChars) {
		m[r] = struct{}{}
	}
	return m
}()

func allowChar(r rune) rune {
	if _, ok := allowedCharsMap[r]; ok {
		return r
	}
	return -1
}

const (
  charWidth = 5
  charHeight = 6
)

func lines(ss ...string) string {
  return strings.Join(ss, "\n")
}
