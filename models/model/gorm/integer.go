package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Int(name, description string) swagger.Definition{
	return model.Property(name, model.Int()).Description(description)
}

func Long(name, description string) swagger.Definition{
	return model.Property(name, model.Long()).Description(description)
}
