package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	posX     float64
	posY     float64
	isYellow bool
}

var mob *ebiten.Image
var mob2 *ebiten.Image
var tiles *ebiten.Image
var gras1 *ebiten.Image
var gras2 *ebiten.Image
var gras3 *ebiten.Image

func (g *Game) Update(screen *ebiten.Image) error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.isYellow = !g.isYellow
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.posY += 5
		if g.isYellow && g.posX > 80 && g.posX < 200 && g.posY > 80 && g.posY < 200 {
			g.posY = 80
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.posY -= 5
		if g.isYellow && g.posX > 80 && g.posX < 200 && g.posY > 80 && g.posY < 200 {
			g.posY = 200
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.posX -= 5
		if g.isYellow && g.posX > 80 && g.posX < 200 && g.posY > 80 && g.posY < 200 {
			g.posX = 200
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.posX += 5
		if g.isYellow && g.posX > 80 && g.posX < 200 && g.posY > 80 && g.posY < 200 {
			g.posX = 80
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})
	for i := 0.0; i < 640.0; i += 16.0 {
		for j := 0.0; j < 480.0; j += 16.0 {
			tilesOp := &ebiten.DrawImageOptions{}
			tilesOp.GeoM.Translate(i, j)
			gt := rand.Int() % 3
			if gt == 0 {
				screen.DrawImage(gras1, tilesOp)
			} else if gt == 1 {
				screen.DrawImage(gras2, tilesOp)
			} else if gt == 2 {
				screen.DrawImage(gras3, tilesOp)
			}
		}
	}

	ebitenutil.DrawRect(screen, 100, 100, 100, 100, color.RGBA{0x00, 0xff, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.posX, g.posY)
	if g.isYellow {
		screen.DrawImage(mob, op)
	} else {
		screen.DrawImage(mob2, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	var err error
	mob, _, err = ebitenutil.NewImageFromFile("mob.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	mob2, _, err = ebitenutil.NewImageFromFile("mob2.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	tiles, _, err = ebitenutil.NewImageFromFile("basictiles.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	sx := 0 * 16
	sy := 8 * 16
	gras1 = tiles.SubImage(image.Rect(sx, sy, sx+16, sy+16)).(*ebiten.Image)
	sx = 1 * 16
	sy = 8 * 16
	gras2 = tiles.SubImage(image.Rect(sx, sy, sx+16, sy+16)).(*ebiten.Image)
	sx = 3 * 16
	sy = 1 * 16
	gras3 = tiles.SubImage(image.Rect(sx, sy, sx+16, sy+16)).(*ebiten.Image)

	game := &Game{
		posX: 50,
		posY: 50,
	}

	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
