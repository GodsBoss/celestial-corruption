package main

import (
  "github.com/GodsBoss/celestial-corruption/pkg/game"

  "github.com/GodsBoss/gggg/pkg/dom"
  "github.com/GodsBoss/gggg/pkg/dominit"
)

func main() {
  win, _ := dom.GlobalWindow()
  doc, _ := win.Document()

  img, _ := doc.CreateImageElement("gfx.png")
  img.On(
    // Success
    func(){
      hint, _ := doc.GetElementByID("hint")
      dom.RemoveNode(hint)

      dominit.Run(game.New(img))
    },

    // Fail
    func(err interface{}){},
  )

  <-make(chan struct{}, 0)
}
