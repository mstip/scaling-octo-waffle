package pkg

import (
	"errors"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
	"image"
	"io/ioutil"
	"log"
)

type Assets struct {
	NormalFont font.Face
	Grass      *ebiten.Image
	RedFlowers *ebiten.Image
	Dude       *ebiten.Image
	Tree       *ebiten.Image
	Stone      *ebiten.Image
	LootBox    *ebiten.Image
}

func (a *Assets) GetAssetById(id int) (*ebiten.Image, error) {
	if id == 0 {
		return a.Grass, nil
	} else if id == 1 {
		return a.RedFlowers, nil
	} else if id == 2 {
		return a.Dude, nil
	} else if id == 3 {
		return a.Tree, nil
	} else if id == 4 {
		return a.Stone, nil
	} else if id == 5 {
		return a.LootBox, nil
	}
	return nil, errors.New("asset ID not found")
}

func NewAssets() (*Assets, error) {
	var a Assets
	dat, err := ioutil.ReadFile("./assets/UbuntuMono-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(dat)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	a.NormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	tiles, _, err := ebitenutil.NewImageFromFile("./assets/roguelikeSheet_transparent.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	sx := 51
	sy := 272
	a.Grass = tiles.SubImage(image.Rect(sx, sy, sx+FieldSize, sy+FieldSize)).(*ebiten.Image)
	sx = 51
	sy = 119
	a.RedFlowers = tiles.SubImage(image.Rect(sx, sy, sx+FieldSize, sy+FieldSize)).(*ebiten.Image)
	sx = 221
	sy = 153
	a.Tree = tiles.SubImage(image.Rect(sx, sy, sx+FieldSize, sy+FieldSize)).(*ebiten.Image)
	sx = 952
	sy = 357
	a.Stone = tiles.SubImage(image.Rect(sx, sy, sx+FieldSize, sy+FieldSize)).(*ebiten.Image)
	sx = 764
	sy = 170
	a.LootBox = tiles.SubImage(image.Rect(sx, sy, sx+FieldSize, sy+FieldSize)).(*ebiten.Image)

	dude, _, err := ebitenutil.NewImageFromFile("./assets/dude1.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	sx = 0
	sy = 8
	a.Dude = dude.SubImage(image.Rect(sx, sy, sx+FieldSize, sy+FieldSize)).(*ebiten.Image)

	return &a, nil
}
