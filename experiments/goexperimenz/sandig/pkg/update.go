package pkg

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"math"
	"os"
)

func (g *Game) Update(_ *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if g.WorldData.GameMode == GameModeNormal {
		newPosX := g.WorldData.GetPlayer().PosX
		newPosY := g.WorldData.GetPlayer().PosY
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
			newPosY -= 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
			newPosY += 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
			newPosX -= 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
			newPosX += 1
		}

		// no  collision -> set new pos
		if obj, _ := g.WorldData.GetObjAt(newPosX, newPosY); obj == nil {
			g.WorldData.GetPlayer().PosX = newPosX
			g.WorldData.GetPlayer().PosY = newPosY
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			mouseX, mouseY := ebiten.CursorPosition()
			g.WorldData.SelectedFieldX = mouseX / FieldSize
			g.WorldData.SelectedFieldY = mouseY / FieldSize
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			obj, i := g.WorldData.GetObjAt(g.WorldData.SelectedFieldX, g.WorldData.SelectedFieldY)
			// check there is an object and its not the player
			if obj != nil && obj.ID != g.WorldData.GetPlayer().ID {
				// check the distance is in an 1er range
				if math.Abs(float64(obj.PosX-g.WorldData.GetPlayer().PosX)) <= 1 &&
					math.Abs(float64(obj.PosY-g.WorldData.GetPlayer().PosY)) <= 1 {
					// check attackable
					if obj.IsAttackable {
						fmt.Println("ATTACK!")
						g.WorldData.Objects[i].HP -= g.WorldData.GetPlayer().Dmg
						fmt.Println("obj.hp", obj.HP, "p dmg ", g.WorldData.GetPlayer().Dmg)
					} else if obj.IsLoot {
						g.WorldData.GameMode = GameModeMenu
						g.WorldData.CurrentMenu = MenuLoot
					}
				}
			}
		}
	} else if g.WorldData.GameMode == GameModeMenu {

	}

	g.WorldData.ObjectActions()

	return nil
}

func AttackAction(source Object, target Object) {

}
