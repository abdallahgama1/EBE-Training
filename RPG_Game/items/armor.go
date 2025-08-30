package items

import (
	"RPG_Game/interfaces"
	"fmt"
)

type Armor struct {
	BaseItem
	Defense int
}

func NewArmor(name string, defense int) *Armor {
	return &Armor{BaseItem{Name: name}, defense}
}

func (a *Armor) Use(target interfaces.Defendable) {
	fmt.Printf("%s equips armor %s and gains +%d defense (not implemented yet).\n", target.GetName(), a.Name, a.Defense)
}
