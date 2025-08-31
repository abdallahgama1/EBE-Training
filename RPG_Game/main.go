package main

import (
	"RPG_Game/characters"
	"RPG_Game/equipment"
	"RPG_Game/items"
	"RPG_Game/monsters"
	"fmt"
)

func main() {
	warrior := characters.NewWarrior("Arthas")
	mage := characters.NewMage("Jaina")
	monster := monsters.NewMonster("Goblin", 60, 10)

	sword := items.NewWeapon("Sword of Dawn", 20)
	armor := items.NewArmor("Steel Armor", 5)
	potion := items.NewPotion("Healing Potion", 30)

	// Equip characters
	equipment.EquipItem(warrior, sword)
	equipment.EquipItem(warrior, armor)

	// Battle
	fmt.Println("\n--- Battle Starts ---")
	warrior.Attack(monster)
	mage.Attack(monster)
	monster.Attack(warrior)

	// Heal warrior
	potion.Use(warrior)

	fmt.Println("\n--- Battle Ends ---")
}
