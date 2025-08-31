package items

import (
	"RPG_Game/interfaces"
	"fmt"
)

type Potion struct {
	BaseItem
	HealAmount int
}

func NewPotion(name string, heal int) *Potion {
	return &Potion{BaseItem{Name: name}, heal}
}

func (p *Potion) Use(target interfaces.Defendable) {
	// Defensive cast to concrete type
	if char, ok := target.(interface {
		GetName() string
		GetHealth() int
		TakeDamage(int)
	}); ok {
		fmt.Printf("%s drinks %s and heals %d HP!\n", char.GetName(), p.Name, p.HealAmount)
		char.TakeDamage(-p.HealAmount) // negative damage = healing
	}
}
