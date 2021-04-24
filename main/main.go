package main

import (
  "github.com/GodsBoss/gggg/pkg/dom"
)

func main() {
  win, _ := dom.GlobalWindow()
  doc, _ := win.Document()

  hint, _ := doc.GetElementByID("hint")
  dom.RemoveNode(hint)

  canvas, _ := doc.CreateCanvasElement()
  canvas.SetSize(800, 600)
  gameElement, _ := doc.GetElementByID("game")
  gameElement.AppendChild(canvas)
}
