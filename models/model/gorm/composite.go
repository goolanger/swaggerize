package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
	"strings"
)

func Composite(name, description string, definition swagger.Definition) swagger.Definition {
	return model.Property(name, definition).Description(description).
		Tag("x-go-custom-tag", "primary_key;auto_increment:false")
}

func CompositeId(definition swagger.Definition) swagger.Definition {
	return model.Property(strings.ToLower(definition.GetName())+"Id", model.Long()).
		Description("id of related model "+definition.GetName()).
		Tag("x-go-custom-tag", "primary_key;auto_increment:false")
}
