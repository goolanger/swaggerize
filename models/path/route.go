package path

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
)

const Inherit = ""

type endpoint struct {
	path, description string
	method            methods.Type

	produces, consumes []mimes.Type
	parameters         []swagger.Parameter
	tags               []swagger.Tag
}

func (r *endpoint) SetPath(s string) swagger.Path {
	r.path = s
	return r
}

func (r *endpoint) SetMethod(m methods.Type) swagger.Path {
	r.method = m
	return r
}

func (r *endpoint) GetMethod() string {
	if r.method == "" {
		r.method = methods.GET
	}
	return string(r.method)
}


func (r *endpoint) GetPath() string {
	return r.path
}

func (r *endpoint) Param(d ...swagger.Parameter) swagger.Path {
	r.parameters = append(r.parameters, d...)
	return r
}

func (r *endpoint) Tag(t ...swagger.Tag) swagger.Path {
	r.tags = append(r.tags, t...)
	return r
}

func (r *endpoint)Produces(p ...mimes.Type) swagger.Path{
	r.produces = append(r.produces, p...)
	return r
}

func (r *endpoint)Consumes(c ...mimes.Type) swagger.Path {
	r.consumes = append(r.consumes, c...)
	return r
}

func (r *endpoint) GetRep() map[string]interface{} {
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

	if len(r.parameters) > 0 {
		var params []map[string]interface{}
		for _, p := range r.parameters {
			params = append(params, p.GetRep())
		}
		rep["parameters"] = params
	}

	return map[string]interface{}{
		r.GetMethod(): rep,
	}
}

func Endpoint(path string) *endpoint {
	return &endpoint{
		method: methods.GET,
		path: path,
	}
}