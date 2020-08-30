package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func String (name, description string) swagger.Definition {
	return model.Property(name, model.String()).Description(description)
}

func File(name, description string) swagger.Definition {
	return model.Property(name, model.File()).Description(description)
}

func Byte(name, description string) swagger.Definition {
	return model.Property(name, model.Byte()).Description(description).
		Tag("x-go-custom-tag", "schema:\"-\"")
}

func TimeStamp(name, description string) swagger.Definition {
	return model.Property(name, model.DateTime()).
		Tag("x-go-custom-tag", "gorm:\"Type:timestamp\"").
		Description(description)
}