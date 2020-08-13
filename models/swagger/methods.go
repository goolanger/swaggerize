package swagger

import (
	"github.com/goolanger/swaggerize/models/document"
)

func (specs *Instance) Info(info document.Info) *Instance {
	specs.info = &info
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
	if specs.definitions == nil {
		specs.definitions = []Definition{d}
	} else {
		specs.definitions = append(specs.definitions, d)
	}
	return d
}

func (specs *Instance) Route(p Path) Path {
	if specs.paths == nil {
		specs.paths = []Path{p}
	} else {
		specs.paths = append(specs.paths, p)
	}
	return p
}

func (specs *Instance) Tag(t Tag) Tag {
	specs.tags = append(specs.tags, t)
	return t
}
