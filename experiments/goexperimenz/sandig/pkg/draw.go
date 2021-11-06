package pkg

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"log"
)

func (g *Game) Draw(screen *ebiten.Image) {
	//TODO: dont pass game
	DrawTiles(screen, g)
	DrawObjs(screen, g)
	DrawSelection(screen, g)

	if g.WorldData.GameMode == GameModeMenu {
		DrawMenu(screen, g.WorldData, g.Assets)
	}

	err := ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"FPS %f TPS %f SX: %d SY: %d",
			ebiten.CurrentFPS(),
			ebiten.CurrentTPS(),
			g.WorldData.SelectedFieldX,
			g.WorldData.SelectedFieldY,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func DrawMenu(screen *ebiten.Image, data *WorldData, assets *Assets) {
	ebitenutil.DrawRect(
		screen,
		50,
		50,
		ScreenWidth-100,
		ScreenHeight-100,
		color.RGBA{R: 0x8b, G: 0x45, B: 0x13, A: 0xd0},
	)
	ebitenutil.DrawRect(
		screen,
		65,
		65,
		ScreenWidth-130,
		ScreenHeight-130,
		color.RGBA{A: 0xd0},
	)
	for i := 0; i < 25; i++ {
		text.Draw(screen, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890!\"§$$$$$$$$", assets.NormalFont, 65, 75+14*i, color.White)
	}
}

func DrawObjs(screen *ebiten.Image, g *Game) {
	for _, obj := range g.WorldData.Objects {
		tilesOp := &ebiten.DrawImageOptions{}
		tilesOp.GeoM.Translate(float64(obj.PosX*FieldSize), float64(obj.PosY*FieldSize))
		asset, err := g.Assets.GetAssetById(obj.AssetId)
		if err != nil {
			log.Fatal(err)
		}
		if err := screen.DrawImage(asset, tilesOp); err != nil {
			log.Fatal(err)
		}
	}
}

func DrawTiles(screen *ebiten.Image, g *Game) {
	for i := 0; i < FieldsCountX; i++ {
		for j := 0; j < FieldsCountY; j++ {
			tilesOp := &ebiten.DrawImageOptions{}
			tilesOp.GeoM.Translate(float64(i*FieldSize), float64(j*FieldSize))
			asset, err := g.Assets.GetAssetById(g.WorldData.Tiles[i][j])
			if err != nil {
				log.Fatal(err)
			}
			if err := screen.DrawImage(asset, tilesOp); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func DrawSelection(screen *ebiten.Image, g *Game) {
	c := color.RGBA{G: 0xff, A: 0xff}
	ebitenutil.DrawLine(
		screen,
		float64(g.WorldData.SelectedFieldX*FieldSize+1),
		float64(g.WorldData.SelectedFieldY*FieldSize),
		float64(g.WorldData.SelectedFieldX*FieldSize+1),
		float64(g.WorldData.SelectedFieldY*FieldSize+FieldSize),
		c,
	)

	ebitenutil.DrawLine(
		screen,
		float64(g.WorldData.SelectedFieldX*FieldSize+FieldSize),
		float64(g.WorldData.SelectedFieldY*FieldSize),
		float64(g.WorldData.SelectedFieldX*FieldSize)+FieldSize,
		float64(g.WorldData.SelectedFieldY*FieldSize+FieldSize),
		c,
	)

	ebitenutil.DrawLine(
		screen,
		float64(g.WorldData.SelectedFieldX*FieldSize),
		float64(g.WorldData.SelectedFieldY*FieldSize),
		float64(g.WorldData.SelectedFieldX*FieldSize+FieldSize),
		float64(g.WorldData.SelectedFieldY*FieldSize),
		c,
	)
	ebitenutil.DrawLine(
		screen,
		float64(g.WorldData.SelectedFieldX*FieldSize),
		float64(g.WorldData.SelectedFieldY*FieldSize+FieldSize-1),
		float64(g.WorldData.SelectedFieldX*FieldSize+FieldSize),
		float64(g.WorldData.SelectedFieldY*FieldSize+FieldSize-1),
		c,
	)
}
