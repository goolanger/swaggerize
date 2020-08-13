package swagger

import (
	"github.com/goolanger/swaggerize/models/document"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
)

type Tag interface {
	GetName() string
}

type Definition interface {
	GetName() string
	GetRep() map[string]interface{}
	GetRef() Definition
}

type Path interface {
	SetPath(string) Path
	GetPath() string
	SetMethod(m methods.Type)Path
	GetMethod() string
	GetRep() map[string]interface{}

	Param(d ...Definition) Path
	Tag(t ...Tag) Path
	Produces(p ...mimes.Type) Path
	Consumes(c ...mimes.Type) Path
}

type Instance struct {
	definitions []Definition
	paths       []Path
	tags        []Tag

	basePath     *string
	externalDocs *document.External
	info         *document.Info
}

func New() *Instance {
	return &Instance{}
}
