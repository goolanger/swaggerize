package obj

import "github.com/goolanger/swaggerize/models/swagger"

type number struct {

}

func (n *number) GetName() string {
	panic("implement me")
}

func (n* number) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"type": "number",
	}
}

func (n *number) GetRef() swagger.Definition {
	return n
}

func Number() *number {
	return &number{}
}

