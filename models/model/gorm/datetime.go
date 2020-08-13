package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func DateTime(name, description string) swagger.Definition {
	return model.Property(name, model.DateTime()).
		Tag("x-go-custom-tag", "gorm:\"Type:timestamp\"").
		Description(description)
}

