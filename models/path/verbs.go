package path

import "github.com/goolanger/swaggerize/models/types/methods"

func Get(path, id string) *endpoint {
	e := Endpoint(path, id)
	e.SetMethod(methods.GET)
	return e
}

func Post(path, id string) *endpoint {
	e := Endpoint(path, id)
	e.SetMethod(methods.POST)
	return e
}

func Put(path, id string) *endpoint {
	e := Endpoint(path, id)
	e.SetMethod(methods.PUT)
	return e
}

func Patch(path, id string) *endpoint {
	e := Endpoint(path, id)
	e.SetMethod(methods.PATCH)
	return e
}

func Delete(path, id string) *endpoint {
	e := Endpoint(path, id)
	e.SetMethod(methods.DELETE)
	return e
}

func Head(path, id string) *endpoint {
	e := Endpoint(path, id)
	e.SetMethod(methods.HEAD)
	return e
}