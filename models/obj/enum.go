package obj

import "github.com/goolanger/swaggerize/models/swagger"

type enum struct {
	def swagger.Definition
	items []interface{}
}

func (e *enum) GetName() string {
	panic("implement me")
}

func (e *enum) GetRep() map[string]interface{} {
	rep := e.def.GetRep()

	if len(e.items)>0{
		rep["enum"] = e.items
	}

	return rep
}

func (e *enum) GetRef() swagger.Definition {
	return e
}

func Enum(definition swagger.Definition, items ...interface{}) *enum {
	return &enum{def: definition, items: items}
}

