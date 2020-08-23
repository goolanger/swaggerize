package model

import "github.com/goolanger/swaggerize/models/swagger"

type double struct {
	format string
}

func (n *double) GetName() string {
	panic("operation not allowed")
}

func (n* double) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"format": n.format,
		"type": "number",
	}
}

func (n *double) GetRef() swagger.Definition {
	return n
}

func Double() *double {
	return &double{format: "double"}
}

func Float() *double {
	return &double{format: "float"}
}
