package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Array(name, description string, definition swagger.Definition) swagger.Definition {
	return model.Property(name, model.Array(definition)).Description(description)
}
