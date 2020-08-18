package path

import (
	"github.com/goolanger/swaggerize/models/security"
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
)

const Inherit = ""

type endpoint struct {
	path, description, id string
	method            methods.Type

	produces, consumes []mimes.Type
	parameters         []swagger.Parameter
	responses          []swagger.Response
	secures			   []swagger.Security
	tags               []swagger.Tag
}

func (e *endpoint) GetId() string {
	return e.id
}

func (e *endpoint) SetId(id string) swagger.Path {
	e.id = id
	return e
}

func (e *endpoint) SetPath(s string) swagger.Path {
	e.path = s
	return e
}

func (e *endpoint) SetMethod(m methods.Type) swagger.Path {
	e.method = m
	return e
}

func (e *endpoint) GetMethod() string {
	if e.method == "" {
		e.method = methods.GET
	}
	return string(e.method)
}


func (e *endpoint) GetPath() string {
	return e.path
}

func (e *endpoint) Params(d ...swagger.Parameter) swagger.Path {
	e.parameters = append(e.parameters, d...)
	return e
}

func (e *endpoint) Secure(c ...swagger.Security) swagger.Path {
	e.secures = append(e.secures, c...)
	return e
}

func (e *endpoint) Tag(t ...swagger.Tag) swagger.Path {
	e.tags = append(e.tags, t...)
	return e
}

func (e *endpoint)Produces(p ...mimes.Type) swagger.Path{
	e.produces = append(e.produces, p...)
	return e
}

func (e *endpoint)Consumes(c ...mimes.Type) swagger.Path {
	e.consumes = append(e.consumes, c...)
	return e
}

func (e *endpoint) Responds(r ...swagger.Response) swagger.Path {
	e.responses = append(e.responses, r...)
	return e
}

func (e *endpoint) GetRep() map[string]interface{} {
	rep:= make(map[string]interface{})

	rep["operationId"] = e.id

	if e.description != "" {
		rep["description"] = e.description
	}

	if len(e.tags) > 0 {
		var tags []string
		for _, t := range e.tags {
			tags = append(tags, t.GetName())
		}
		rep["tags"] = tags
	}

	if len(e.produces) > 0 {
		rep["produces"] = uniq(e.produces)
	}

	if len(e.consumes) > 0 {
		rep["consumes"] = uniq(e.consumes)
	}

	var secures []interface{}
	for _, sec := range e.secures {
		if sec == security.None() {
			secures = []interface{}{}
			rep["security"] = secures
			break
		}
		secures = append(secures, sec.GetRep())
	}
	if len(secures) > 0 {
		rep["security"] = secures
	}


	if len(e.responses) > 0 {
		resp := make(map[string]interface{})

		for _, r := range e.responses {
			if _, ok := resp[r.GetCode()]; !ok {
				resp[r.GetCode()] = r.GetRep()
			}
		}

		rep["responses"] = resp
	}

	if len(e.parameters) > 0 {
		var params []map[string]interface{}
		for _, p := range e.parameters {
			params = append(params, p.GetRep())
		}
		rep["parameters"] = params
	}

	return map[string]interface{}{
		e.GetMethod(): rep,
	}
}

func Endpoint(path, id string) *endpoint {
	return &endpoint{
		method: methods.GET,
		path:   path,
		id:     id,
	}
}

func (e *endpoint) Description(s string) *endpoint {
	e.description = s
	return e
}

func uniq (slice[]mimes.Type) []mimes.Type {
	var result []mimes.Type
	contains := make(map[mimes.Type]bool)

	for _, e := range slice {
		if _, ok := contains[e]; !ok {
			result = append(result, e)
			contains[e] = true
		}
	}
	return result
}