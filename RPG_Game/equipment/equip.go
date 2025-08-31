package equipment

import (
	"RPG_Game/interfaces"
	"fmt"
)

func EquipItem(character interfaces.Character, item interfaces.Item) {
	fmt.Printf("%s equips %s.\n", character.GetName(), item.GetName())
	character.Equip(item)
}
