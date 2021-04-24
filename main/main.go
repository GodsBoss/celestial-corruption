package main

import (
  "github.com/GodsBoss/gggg/pkg/dom"
)

func main() {
  win, _ := dom.GlobalWindow()
  doc, _ := win.Document()
  hint, _ := doc.GetElementByID("hint")
  dom.RemoveNode(hint)
}
