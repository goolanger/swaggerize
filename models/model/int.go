package model

import "github.com/goolanger/swaggerize/models/swagger"

type integer struct {
	format string
}

func (i *integer) GetName() string {
	panic("implement me")
}

func (i *integer) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"format": i.format,
		"type":   "integer",
	}
}

func (i *integer) GetRef() swagger.Definition {
	return i
}

func Int() *integer {
	return &integer{ "int32"}
}

func Int64() *integer {
	return &integer{ "int64"}
}
