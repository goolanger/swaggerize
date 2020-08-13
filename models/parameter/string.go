package params

import (
	"github.com/goolanger/swaggerize/models/types/locations"
)

type stringParam struct {
	name, description *string
	in                locations.Type
	required          *bool
}

func (p *stringParam) GetRep() map[string]interface{} {
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

	var _type string

	if p.in == locations.BODY {
		_type = "schema"
	} else {
		_type = "type"
	}

	if p.in == locations.PATH {
		p.Required(true)
	}

	rep[_type] = "string"

	if p.required != nil {
		rep["required"] = p.required
	}
	return rep
}

func String(name string) *stringParam {
	return &stringParam{
		name: &name,
	}
}

func (p *stringParam) In(location locations.Type) *stringParam {
	p.in = location
	return p
}

func (p *stringParam) Description(description string) *stringParam {
	p.description = &description
	return p
}

func (p *stringParam) Required(required bool) *stringParam {
	p.required = &required
	return p
}
