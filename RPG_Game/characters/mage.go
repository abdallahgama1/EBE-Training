package characters

import (
	"RPG_Game/interfaces"
	"fmt"
)

type Mage struct {
	Character
	Mana int
}

func NewMage(name string) *Mage {
	return &Mage{
		Character: Character{Name: name, Health: 80, Level: 1},
		Mana:      100,
	}
}

func (m *Mage) Attack(target interfaces.Defendable) {
	if m.Mana >= 20 {
		fmt.Printf("%s casts Fireball on %s!\n", m.Name, target.GetName())
		target.TakeDamage(25)
		m.Mana -= 20
	} else {
		fmt.Printf("%s is out of mana and uses staff!\n", m.Name)
		target.TakeDamage(10)
	}
}

func (m *Mage) Equip(item interfaces.Item) {
	fmt.Printf("%s equips %s.\n", m.Name, item.GetName())
}
