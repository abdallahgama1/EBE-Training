package characters

import "fmt"

type Character struct {
	Name   string
	Health int
	Level  int
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) GetHealth() int {
	return c.Health
}

func (c *Character) TakeDamage(damage int) {
	c.Health -= damage
	if c.Health < 0 {
		c.Health = 0
	}
	fmt.Printf("%s takes %d damage! Remaining HP: %d\n", c.Name, damage, c.Health)
}
