package pkg

type Object struct {
	ID           int
	Name         string
	PosX         int
	PosY         int
	AssetId      int
	IsAttackable bool
	HP           int
	Dmg          int
	IsLoot       bool
}

var objIdCounter int

func GetNextObjId() int {
	objIdCounter += 1
	return objIdCounter
}

func NewPlayer() Object {
	return Object{
		ID:           GetNextObjId(),
		AssetId:      2,
		IsAttackable: true,
		HP:           10,
		Dmg:          1,
	}
}

func NewTree() Object {
	return Object{
		ID:           GetNextObjId(),
		AssetId:      3,
		IsAttackable: true,
		HP:           3,
		Dmg:          0,
	}
}

func NewStone() Object {
	return Object{
		ID:           GetNextObjId(),
		AssetId:      4,
		IsAttackable: true,
		HP:           10,
		Dmg:          0,
	}
}

func NewLootBox() Object {
	return Object{
		ID:      GetNextObjId(),
		AssetId: 5,
		HP:      1,
		Dmg:     0,
		IsLoot:  true,
	}
}
