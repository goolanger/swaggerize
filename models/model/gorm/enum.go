package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Enum(name, description string, value swagger.Definition, elements ...interface{}) swagger.Definition {
	return model.Property(name, model.Enum(value, elements...)).Description(description)
}
