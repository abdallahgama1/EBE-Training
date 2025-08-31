package interfaces

type Character interface {
	GetName() string
	GetHealth() int
	Attack(target Defendable)
	Equip(item Item)
}
