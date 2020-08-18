package path

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
)

type scoped struct {
	path, id string

	routes []swagger.Path

	tags   []swagger.Tag
	params []swagger.Parameter

	produces, consumes []mimes.Type
	secures            []swagger.Security
	responses          []swagger.Response
}

func (s *scoped) GetId() string {
	return s.id
}

func (s *scoped) SetId(id string) swagger.Path {
	s.id = id
	return s
}

func (s *scoped) Responds(r ...swagger.Response) swagger.Path {
	s.responses = append(s.responses, r...)
	return s
}

func (s *scoped) SetPath(p string) swagger.Path {
	s.path = p
	return s
}

func (s *scoped) GetPath() string {
	return s.path
}

func (s *scoped) SetMethod(m methods.Type) swagger.Path {
	return s
}

func (s *scoped) GetMethod() string {
	return ""
}

func (s *scoped) GetRep() map[string]interface{} {
	for _, r := range s.routes {
		r.SetId(s.GetId() + r.GetId())
		r.SetPath(s.GetPath() + r.GetPath())
		r.Params(s.params...)
		r.Tag(s.tags...)
		r.Secure(s.secures...)
		r.Produces(s.produces...)
		r.Consumes(s.consumes...)
		r.Responds(s.responses...)
	}
	return nil
}

func (s *scoped) Params(p ...swagger.Parameter) swagger.Path {
	s.params = append(s.params, p...)
	return s
}

func (s *scoped) Tag(t ...swagger.Tag) swagger.Path {
	s.tags = append(s.tags, t...)
	return s
}

func (s *scoped) Produces(p ...mimes.Type) swagger.Path {
	s.produces = append(s.produces, p...)
	return s
}

func (s *scoped) Consumes(c ...mimes.Type) swagger.Path {
	s.consumes  = append(s.consumes, c...)
	return s
}

func (s *scoped) Secure(c ...swagger.Security) swagger.Path {
	s.secures = append(s.secures, c...)
	return s
}

func (s *scoped) Routes(p ...swagger.Path) *scoped {
	s.routes = append(s.routes, p...)
	return s
}

func Scope(path, id string) *scoped {
	return &scoped{
		path: path,
		id:id,
	}
}



