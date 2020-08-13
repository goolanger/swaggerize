package obj

import "github.com/goolanger/swaggerize/models/swagger"

type integer struct {

}

func (i integer) GetName() string {
	panic("implement me")
}

func (i integer) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"format": "int64",
		"type":   "integer",
	}
}

func (i integer) GetRef() swagger.Definition {
	return i
}

func Int() *integer {
	return &integer{}
}
