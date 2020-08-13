package params

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/locations"
)

type param struct {
	name, description *string
	in                locations.Type
	required          *bool
	swagger.Definition
}

func (p *param) GetRep() map[string]interface{} {
	rep := make(map[string]interface{})

	if p.in != "" {
		rep["in"] = p.in
	}
	if p.name != nil {
		rep["name"] = p.name
	}
	if p.description != nil {
		rep["description"] = p.description
	}
	if p.Definition != nil {
		var _type string

		if p.in == locations.BODY {
			_type = "schema"
		} else {
			_type = "type"
		}

		if p.in == locations.PATH {
			p.Required(true)
		}

		rep[_type] = p.Definition.GetRef().GetRep()
	}
	if p.required != nil {
		rep["required"] = p.required
	}
	return rep
}

func Param(name string, definition swagger.Definition) *param {
	return &param{
		Definition:definition,
		name: &name,
	}
}

func (p *param) In(location locations.Type) *param {
	p.in = location
	return p
}

func (p *param) Description(description string) *param {
	p.description = &description
	return p
}

func (p *param) Required(required bool) *param {
	p.required = &required
	return p
}