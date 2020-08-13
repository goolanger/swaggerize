package model

import "github.com/goolanger/swaggerize/models/swagger"

type stringModel struct {

}

func (s stringModel) GetName() string {
	panic("operation not allowed")
}

func (s stringModel) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"type":"string",
	}
}

func (s stringModel) GetRef() swagger.Definition {
	return s
}

func String() *stringModel {
	return &stringModel{}
}

