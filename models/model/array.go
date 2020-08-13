package model

import "github.com/goolanger/swaggerize/models/swagger"

type array struct {
	items swagger.Definition
}

func (a *array) GetName() string {
	panic("operation not allowed")
}

func (a *array) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"type":  "array",
		"items": a.items.GetRep(),
	}
}

func (a *array) GetRef() swagger.Definition {
	return a
}

func Array(items swagger.Definition) *array {
	return &array{items: items}
}

