package characters

import (
	"RPG_Game/interfaces"
	"fmt"
)

type Warrior struct {
	Character
	Strength int
}

func NewWarrior(name string) *Warrior {
	return &Warrior{
		Character: Character{Name: name, Health: 120, Level: 1},
		Strength:  15,
	}
}

func (w *Warrior) Attack(target interfaces.Defendable) {
	fmt.Printf("%s slashes at %s!\n", w.Name, target.GetName())
	target.TakeDamage(w.Strength)
}

func (w *Warrior) Equip(item interfaces.Item) {
	fmt.Printf("%s equips %s.\n", w.Name, item.GetName())
}
