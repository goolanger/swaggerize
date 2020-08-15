package model

import "github.com/goolanger/swaggerize/models/swagger"

type date struct {

}

func (d *date) GetName() string {
	panic("implement me")
}

func (d *date) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"format": "date",
		"type": "string",
	}
}

func (d *date) GetRef() swagger.Definition {
	return d
}

func Date() *date {
	return &date{}
}


