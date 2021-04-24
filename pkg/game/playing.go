package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
  passed int
  spriteFactory *spriteFactory
}

var _ state = &playing{}

func (p *playing) init() {
  p.passed = 0
}

func (p *playing) tick(ms int)  (next string) {
  p.passed += ms

  if p.passed > 1000 {
    return "title"
  }

  return ""
}

func (p *playing) receiveKeyEvent(event interaction.KeyEvent) (next string){
  return ""
}

func (p *playing) renderable() renderable {
  return p.spriteFactory.create("bg_playing", 0, 0, 0)
}
