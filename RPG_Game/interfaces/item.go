package interfaces

type Item interface {
	Use(target Defendable)
	GetName() string
}
