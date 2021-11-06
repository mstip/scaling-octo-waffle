package pkg

type WorldData struct {
	Tiles          [FieldsCountX][FieldsCountY]int
	Objects        []Object
	SelectedFieldX int
	SelectedFieldY int
	GameMode       int
	CurrentMenu    int
}

const GameModeNormal = 0
const GameModeMenu = 1
const GameModeSettings = 2

const MenuLoot = 0

func NewWorld() (*WorldData, error) {
	var w WorldData
	w.GameMode = GameModeMenu
	w.Tiles[5][5] = 1
	w.Tiles[0][0] = 1
	w.Tiles[FieldsCountX-1][0] = 1
	w.Tiles[FieldsCountX-1][FieldsCountY-1] = 1
	w.Tiles[0][FieldsCountY-1] = 1

	player := NewPlayer()
	player.PosX = 10
	player.PosY = 10
	w.Objects = append(w.Objects, player)

	tree := NewTree()
	tree.PosX = 10
	tree.PosY = 12
	w.Objects = append(w.Objects, tree)

	stone := NewStone()
	stone.PosX = 11
	stone.PosY = 12
	w.Objects = append(w.Objects, stone)

	return &w, nil
}

func (w *WorldData) GetObjAt(x, y int) (*Object, int) {
	for i, obj := range w.Objects {
		if obj.PosX == x && obj.PosY == y {
			return &obj, i
		}
	}
	return nil, 0
}

func (w *WorldData) GetPlayer() *Object {
	return &w.Objects[0]
}

func (w *WorldData) ObjectActions() {
	var newObjects []Object
	for _, obj := range w.Objects {
		// dead
		if obj.HP <= 0 {
			lootBox := NewLootBox()
			lootBox.PosX = obj.PosX
			lootBox.PosY = obj.PosY
			newObjects = append(newObjects, lootBox)
			continue
		}
		newObjects = append(newObjects, obj)
	}
	w.Objects = newObjects
}
