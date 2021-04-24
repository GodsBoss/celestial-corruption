package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {}

var _ state = &playing{}

func (p *playing) init() {}

func (p *playing) tick(ms int)  (next string) {
  return ""
}

func (p *playing) receiveKeyEvent(event interaction.KeyEvent) (next string){
  return ""
}

func (p *playing) sprites() map[string]sprite {
  return map[string]sprite{}
}
