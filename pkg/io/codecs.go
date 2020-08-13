package io

import (
	"github.com/go-yaml/yaml"
	"github.com/goolanger/swaggerize/models/swagger"
)

func Encode(swagger *swagger.Instance) ([]byte, error) {
	return yaml.Marshal(swagger.Encode())
}