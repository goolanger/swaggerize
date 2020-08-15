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

func Response(code int, desctiption string) *respond {
	return &respond{
		description: desctiption,
		code: code,
	}
}

func (r *respond) Schema(d swagger.Definition) *respond {
	r.def = d
	return r
}

