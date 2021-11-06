package pkg

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

const ScreenWidth = 640
const ScreenHeight = 480
const FieldSize = 16
const FieldsCountX = ScreenWidth / FieldSize
const FieldsCountY = ScreenHeight / FieldSize

type Game struct {
	WorldData *WorldData
	Assets    *Assets
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Run() {
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Sandig")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func NewGame() (*Game, error) {
	var err error
	var g Game

	g.WorldData, err = NewWorld()
	if err != nil {
		log.Fatal(err)
	}

	g.Assets, err = NewAssets()
	if err != nil {
		log.Fatal(err)
	}

	return &g, nil
}
