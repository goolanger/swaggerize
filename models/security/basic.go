package security

import "github.com/goolanger/swaggerize/models/swagger"

type basic struct {
	name        string
	description *string
}

func (b *basic) GetName() string {
	return b.name
}

func (b *basic) GetRep() map[string]interface{} {
	rep := map[string]interface{} {
		"type": "basic",
	}

	if b.description != nil {
		rep["description"] = b.description
	}

	return rep
}

func (b *basic) GetRef() swagger.Security {
	return Reference(b.GetName())
}

func Basic(name string) *basic {
	return &basic{name: name}
}

func (b *basic) Description(d string) *basic {
	b.description = &d
	return b
}

