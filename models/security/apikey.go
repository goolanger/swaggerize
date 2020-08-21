package security

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/locations"
)

type apikey struct {
	name  string
	description, header *string
	in *locations.Type
}

func (a *apikey) GetName() string {
	return a.name
}

func (a *apikey) GetRep() map[string]interface{} {
	rep := map[string]interface{}{
		"type": "apiKey",
	}

	if a.header == nil {
		rep["name"] = a.GetName()
	} else {
		rep["name"] = a.header
	}

	if a.in == nil {
		rep["in"] = locations.HEADER
	} else {
		rep["in"] = a.in
	}

	if a.description != nil {
		rep["description"] = a.description
	}

	return rep
}

func (a *apikey) GetRef() swagger.Security {
	return Reference(a.GetName())
}

func ApiKey(name string) *apikey {
	return &apikey{name: name}
}

func (a *apikey) In (p locations.Type) *apikey {
	a.in = &p
	return a
}

func (a *apikey) Description (p string) *apikey {
	a.description = &p
	return a
}

func (a *apikey) Header (h string) *apikey {
	a.header = &h
	return a
}
