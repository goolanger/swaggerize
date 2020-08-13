package model

import "github.com/goolanger/swaggerize/models/swagger"

type file struct {

}

func (f *file) GetName() string {
	panic("invalid operation exception")
}

func (f *file) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"type":   "string",
		"format": "byte",
	}
}

func (f *file) GetRef() swagger.Definition {
	return f
}

func File() *file {
	return &file{}
}
