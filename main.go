package main

import (
	"log"

	"github.com/Calvinsd/mygames/runnergame"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := runnergame.Init()
	ebiten.SetWindowSize(game.GetWidth()*2, game.GetHeight()*2)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
