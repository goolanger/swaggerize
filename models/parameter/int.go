package params

import (
	"github.com/goolanger/swaggerize/models/types/locations"
)

type intParam struct {
	name, description *string
	in                locations.Type
	required          *bool
}

func (p *intParam) GetRep() map[string]interface{} {
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

	rep[_type] = "integer"

	if p.required != nil {
		rep["required"] = p.required
	}
	return rep
}

func Int(name string) *intParam {
	return &intParam{
		name: &name,
	}
}

func (p *intParam) In(location locations.Type) *intParam {
	p.in = location
	return p
}

func (p *intParam) Description(description string) *intParam {
	p.description = &description
	return p
}

func (p *intParam) Required(required bool) *intParam {
	p.required = &required
	return p
}
