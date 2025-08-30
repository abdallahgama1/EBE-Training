package interfaces

type Attackable interface {
	Attack(target Defendable)
}

type Defendable interface {
	TakeDamage(damage int)
	GetName() string
	GetHealth() int
}
