package security

import (
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/locations"
)

type apiKey struct {
	name             string
	description, key *string
	in               *locations.Type
}

func (a *apiKey) GetName() string {
	return a.name
}

func (a *apiKey) GetRep() map[string]interface{} {
	rep := map[string]interface{}{
		"type": "apiKey",
	}

	if a.key == nil {
		rep["name"] = a.GetName()
	} else {
		rep["name"] = a.key
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

func (a *apiKey) GetRef() swagger.Security {
	return Reference(a.GetName())
}

func ApiKey(name string) *apiKey {
	return &apiKey{name: name}
}

func (a *apiKey) In (p locations.Type) *apiKey {
	a.in = &p
	return a
}

func (a *apiKey) Description (p string) *apiKey {
	a.description = &p
	return a
}

func (a *apiKey) Key(h string) *apiKey {
	a.key = &h
	return a
}
