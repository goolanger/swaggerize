package model

import "github.com/goolanger/swaggerize/models/swagger"

type property struct {
	name, description *string
	tags map[string]interface{}
	def swagger.Definition
}

func (p *property) GetName() string {
	return *p.name
}

func (p *property) GetRep() map[string]interface{} {
	rep := p.def.GetRep()

	for k, v := range p.tags {
		rep[k] = v
	}

	if p.description != nil {
		rep["description"] = p.description
	}

	return rep
}

func (p *property) GetRef() swagger.Definition {
	return p
}

func Property(name string, definition swagger.Definition) *property {
	return &property{name: &name, def: definition}
}

func (p *property) Tag(key, value string) *property {
	if p.tags == nil {
		p.tags = make(map[string]interface{})
	}
	p.tags[key] = value
	return p
}

func (p *property) Description(description string) *property {
	p.description = &description
	return p
}

