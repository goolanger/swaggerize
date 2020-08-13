package response

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"strconv"
)

type respond struct {
	code int
	description string
	def swagger.Definition
}

func (r *respond) GetCode() string {
	return strconv.Itoa(r.code)
}

func (r *respond) GetRep() map[string]interface{} {
	rep := make(map[string]interface{})

	rep["description"] = r.description

	if r.def != nil {
		rep["schema"] = r.def.GetRep()
	}

	return rep
}

func Response(code int, definition swagger.Definition) *respond {
	return &respond{
		def:definition,
		code: code,
	}
}

func (r *respond) Description(d string) *respond {
	r.description = d
	return r
}

