package gorm

import (
	"fmt"
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Translate(name, description string) swagger.Definition {
	return model.Property(name, model.String()).
		Tag("x-go-custom-tag", fmt.Sprintf("translate:\"%s\"", name)).
		Description(description)
}
