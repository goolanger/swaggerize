package obj

import "github.com/goolanger/swaggerize/models/swagger"

type _array struct {
	items swagger.Definition
}

func (a *_array) GetName() string {
	panic("operation not allowed")
}

func (a *_array) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"type":  "array",
		"items": a.items.GetRep(),
	}
}

func (a *_array) GetRef() swagger.Definition {
	return a
}

func Array(items swagger.Definition) *_array {
	return &_array{items: items}
}

