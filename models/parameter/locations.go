package params

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/locations"
)

func Form(name string, definition swagger.Definition) *param {
	return Param(name, definition).In(locations.FORM)
}

func Path(name string, definition swagger.Definition) *param {
	return Param(name, definition).In(locations.PATH)
}

func Body(name string, definition swagger.Definition) *param {
	return Param(name, definition).In(locations.BODY)
}

func Header(name string, definition swagger.Definition) *param {
	return Param(name, definition).In(locations.HEADER)
}

func Query(name string, definition swagger.Definition) *param {
	return Param(name, definition).In(locations.QUERY)
}