package model

import "github.com/goolanger/swaggerize/models/swagger"

type stringModel struct {
	format *string
}

func (s *stringModel) GetName() string {
	panic("operation not allowed")
}

func (s *stringModel) GetRep() map[string]interface{} {
	rep := map[string]interface{}{
		"type":"string",
	}
	if s.format != nil {
		rep["format"] = s.format
	}
	return rep
}

func (s *stringModel) GetRef() swagger.Definition {
	return s
}

func String() *stringModel {
	return &stringModel{}
}

func Binary() *stringModel {
	var format = "binary"
	return &stringModel{format: &format}
}

func Byte() *stringModel {
	var format = "byte"
	return &stringModel{format: &format}
}

func Date() *stringModel {
	var format = "date"
	return &stringModel{format: &format}
}

func DateTime() *stringModel {
	var format = "date-time"
	return &stringModel{format: &format}
}

func File() *stringModel {
	var format = "file"
	return &stringModel{format: &format}
}

func Password() *stringModel {
	var format = "password"
	return &stringModel{format: &format}
}


