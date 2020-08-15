package security

import "github.com/goolanger/swaggerize/models/swagger"

type reference struct {
	name string
	values []interface{}
}

func (r *reference) GetName() string {
	return r.name
}

func (r *reference) GetRep() map[string]interface{} {
	return map[string]interface{}{
		r.GetName(): r.values,
	}
}

func (r *reference) GetRef() swagger.Security {
	return r
}

func Reference(name string, values ...interface{}) *reference {
	return &reference{name, values}
}

