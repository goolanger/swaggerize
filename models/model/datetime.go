package model

import "github.com/goolanger/swaggerize/models/swagger"

type datetime struct {

}

func (d *datetime) GetName() string {
	panic("implement me")
}

func (d *datetime) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"format": "date-time",
		"type": "string",
	}
}

func (d *datetime) GetRef() swagger.Definition {
	return d
}

func DateTime() *datetime {
	return &datetime{}
}

