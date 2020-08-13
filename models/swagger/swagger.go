package swagger

import (
	"github.com/goolanger/swaggerize/models/document"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
	"github.com/goolanger/swaggerize/models/types/scheme"
)

type Tag interface {
	GetName() string
	GetRep()map[string]interface{}
}

type Definition interface {
	GetName() string
	GetRep() map[string]interface{}
	GetRef() Definition
}

type Parameter interface {
	GetRep()map[string]interface{}
}

type Path interface {
	SetPath(string) Path
	GetPath() string
	SetMethod(m methods.Type)Path
	GetMethod() string
	GetRep() map[string]interface{}


	Consumes(c ...mimes.Type) Path
	Param(d ...Parameter) Path
	Produces(p ...mimes.Type) Path
	Responds(r ...Response) Path
 	Tag(t ...Tag) Path

}

type Response interface {
	GetCode() string
	GetRep() map[string]interface{}
}

type Instance struct {
	definitions []Definition
	paths       []Path
	tags        []Tag

	basePath, host *string
	schemes        []scheme.Type
	externalDocs   *document.External
	info           *document.Info
}

func New() *Instance {
	return &Instance{}
}

func (specs *Instance) Info(info document.Info) *Instance {
	specs.info = &info
	return specs
}

func (specs *Instance) Schemes(s ...scheme.Type) *Instance {
	specs.schemes = append(specs.schemes, s...)
	return specs
}

func (specs *Instance) Host(h string) *Instance {
	specs.host = &h
	return specs
}

func (specs *Instance) BasePath(base string) *Instance {
	specs.basePath = &base
	return specs
}

func (specs *Instance) ExternalDocs(docs document.External) *Instance {
	specs.externalDocs = &docs
	return specs
}

func (specs *Instance) Define(d Definition) Definition {
	specs.definitions = append(specs.definitions, d)
	return d
}

func (specs *Instance) Route(p Path) Path {
	specs.paths = append(specs.paths, p)
	return p
}

func (specs *Instance) Tag(t Tag) Tag {
	specs.tags = append(specs.tags, t)
	return t
}