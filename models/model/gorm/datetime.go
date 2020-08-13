package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

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

func DateTime(name, description string) swagger.Definition {
	return model.Property(name, &datetime{}).
		Tag("x-go-custom-tag", "gorm:\"Type:timestamp\"").
		Description(description)
}

