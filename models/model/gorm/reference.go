package gorm

import (
	"fmt"
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
	"strings"
)

func Reference(definition swagger.Definition) swagger.Definition {
	return model.Property(strings.ToLower(definition.GetName()) + "ReferenceId", model.Long()).
		Description(fmt.Sprintf("%s model reference", definition.GetName()))
}
