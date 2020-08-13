package path

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
)

const Inherit = ""

type Route struct {
	path, description string
	method            methods.Type

	produces, consumes []mimes.Type
	parameters         []swagger.Definition
	tags               []swagger.Tag
}

func (r *Route) SetPath(s string) swagger.Path {
	r.path = s
	return r
}

func (r *Route) SetMethod(m methods.Type) swagger.Path {
	r.method = m
	return r
}

func (r *Route) GetMethod() string {
	if r.method == "" {
		r.method = methods.GET
	}
	return string(r.method)
}


func (r *Route) GetPath() string {
	return r.path
}

func (r *Route) Param(d ...swagger.Definition) swagger.Path {
	r.parameters = append(r.parameters, d...)
	return r
}

func (r *Route) Tag(t ...swagger.Tag) swagger.Path {
	r.tags = append(r.tags, t...)
	return r
}

func (r *Route)Produces(p ...mimes.Type) swagger.Path{
	r.produces = append(r.produces, p...)
	return r
}

func (r *Route)Consumes(c ...mimes.Type) swagger.Path {
	r.consumes = append(r.consumes, c...)
	return r
}

func (r *Route) GetRep() map[string]interface{} {
	rep:= make(map[string]interface{})

	if r.description != "" {
		rep["description"] = r.description
	}

	if len(r.tags) > 0 {
		var tags []string
		for _, t := range r.tags {
			tags = append(tags, t.GetName())
		}
		rep["tags"] = tags
	}

	if len(r.produces) > 0 {
		rep["produces"] = r.produces
	}

	if len(r.consumes) > 0 {
		rep["consumes"] = r.consumes
	}

	return map[string]interface{}{
		r.GetMethod(): rep,
	}
}

func NewRoute(path string) *Route {
	return &Route{
		method: methods.GET,
		path: path,
	}
}