package obj

import "github.com/goolanger/swaggerize/models/swagger"

type _string struct {

}

func (s _string) GetName() string {
	panic("operation not allowed")
}

func (s _string) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"type":"string",
	}
}

func (s _string) GetRef() swagger.Definition {
	return s
}

func String() *_string {
	return &_string{}
}

