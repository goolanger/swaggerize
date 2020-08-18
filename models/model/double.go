package model

import "github.com/goolanger/swaggerize/models/swagger"

type double struct {

}

func (n *double) GetName() string {
	panic("implement me")
}

func (n* double) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"format": "double",
		"type": "number",
	}
}

func (n *double) GetRef() swagger.Definition {
	return n
}

func Double() *double {
	return &double{}
}

