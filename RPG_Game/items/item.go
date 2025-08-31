package items

import "RPG_Game/interfaces"

type BaseItem struct {
	Name string
}

func (i *BaseItem) GetName() string {
	return i.Name
}

// Default implementation (can be overridden)
func (i *BaseItem) Use(target interfaces.Defendable) {}
