package model

import "github.com/goolanger/swaggerize/models/swagger"

type reference struct {
	name, ref string
}

func (r *reference) GetName() string {
	return r.name
}

func (r *reference) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"$ref": r.ref,
	}
}

func (r *reference) GetRef() swagger.Definition {
	return r
}

func Reference(name, ref string) *reference {
	return &reference{name, ref}
}

