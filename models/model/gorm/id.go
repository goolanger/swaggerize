package gorm

import (
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Id() swagger.Definition {
	return model.Property("id", model.Int64() ).
		Tag("x-go-custom-tag", "gorm:\"primary_key\"")
}

