package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Boolean(name, description string) swagger.Definition{
	return model.Property(name, model.Boolean()).Description(description)
}
