package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Float(name, description string) swagger.Definition {
	return model.Property(name, model.Float()).Description(description)
}

func Double(name, description string) swagger.Definition{
	return model.Property(name, model.Double()).Description(description)
}
