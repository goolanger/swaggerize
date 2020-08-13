package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Index(name, description string, definition swagger.Definition) swagger.Definition {
	return model.Property(name, definition).
		Tag("x-go-custom-tag", "gorm:\"unique_index\"").
		Description(description)
}