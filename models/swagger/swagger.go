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

type Security interface {
	GetName() string
	GetRep() map[string]interface{}
	GetRef() Security
}

type Path interface {
	SetPath(string) Path
	GetPath() string
	SetMethod(m methods.Type)Path
	GetMethod() string
	GetRep() map[string]interface{}
	GetId() string
	SetId(id string) Path


	Consumes(c ...mimes.Type) Path
	Params(d ...Parameter) Path
	Produces(p ...mimes.Type) Path
	Responds(r ...Response) Path
	Secure(c ...Security) Path
 	Tag(t ...Tag) Path
}

type Response interface {
	GetCode() string
	GetRep() map[string]interface{}
}

type Instance struct {
	definitions         []Definition
	paths               []Path
	securities, secures []Security
	tags                []Tag

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

func (specs *Instance) Host(host string) *Instance {
	specs.host = &host
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

func (specs *Instance) Security(security Security) Security {
	specs.securities = append(specs.securities, security)
	return security
}

func (specs *Instance) Secure(security Security) Security {
	specs.secures = append(specs.secures, security)
	return security
}

func (specs *Instance) Tag(t Tag) Tag {
	for _, t1 := range specs.tags {
		if t1.GetName() == t.GetName() {
			return t1
		}
	}
	specs.tags = append(specs.tags, t)
	return t
}