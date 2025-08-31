package monsters

import (
	"RPG_Game/interfaces"
	"fmt"
)

type Monster struct {
	Name   string
	Health int
	Power  int
}

func NewMonster(name string, health, power int) *Monster {
	return &Monster{Name: name, Health: health, Power: power}
}

func (m *Monster) GetName() string {
	return m.Name
}

func (m *Monster) GetHealth() int {
	return m.Health
}

func (m *Monster) TakeDamage(damage int) {
	m.Health -= damage
	if m.Health < 0 {
		m.Health = 0
	}
	fmt.Printf("%s takes %d damage! Remaining HP: %d\n", m.Name, damage, m.Health)
}

func (m *Monster) Attack(target interfaces.Defendable) {
	fmt.Printf("%s attacks %s!\n", m.Name, target.GetName())
	target.TakeDamage(m.Power)
}
