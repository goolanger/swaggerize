package path

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
)

type Scope struct {
	path string

	routes []swagger.Path

	tags   []swagger.Tag
	params []swagger.Definition

	produces, consumes []mimes.Type
}

func (s *Scope) SetPath(p string) swagger.Path {
	s.path = p
	return s
}

func (s *Scope) SetMethod(m methods.Type) swagger.Path {
	return s
}

func (s *Scope) GetMethod() string {
	return ""
}

func (s *Scope) GetPath() string {
	return s.path
}

func (s *Scope) GetRep() map[string]interface{} {
	for _, r := range s.routes {
		r.SetPath(s.GetPath() + r.GetPath())
		r.Param(s.params...)
		r.Tag(s.tags...)
		r.Produces(s.produces...)
		r.Consumes(s.consumes...)
	}
	return nil
}

func (s *Scope) Param(p ...swagger.Definition) swagger.Path {
	s.params = append(s.params, p...)
	return s
}

func (s *Scope) Tag(t ...swagger.Tag) swagger.Path {
	s.tags = append(s.tags, t...)
	return s
}

func (s *Scope) Produces(p ...mimes.Type) swagger.Path {
	s.produces = append(s.produces, p...)
	return s
}

func (s *Scope) Consumes(c ...mimes.Type) swagger.Path {
	s.consumes  = append(s.consumes, c...)
	return s
}

func (s *Scope) Route(p ...swagger.Path) *Scope {
	s.routes = append(s.routes, p...)
	return s
}

func NewScope(path string) *Scope {
	return &Scope{
		path: path,
	}
}



