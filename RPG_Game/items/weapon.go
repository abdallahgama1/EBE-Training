package items

import (
	"RPG_Game/interfaces"
	"fmt"
)

type Weapon struct {
	BaseItem
	Damage int
}

func NewWeapon(name string, damage int) *Weapon {
	return &Weapon{BaseItem{Name: name}, damage}
}

func (w *Weapon) Use(target interfaces.Defendable) {
	fmt.Printf("%s is struck with weapon %s for %d damage!\n", target.GetName(), w.Name, w.Damage)
	target.TakeDamage(w.Damage)
}
