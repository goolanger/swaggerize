package obj

import "github.com/goolanger/swaggerize/models/swagger"

type boolean struct {

}

func (b *boolean) GetName() string {
	panic("implement me")
}

func (b *boolean) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"type": "boolean",
	}
}

func (b *boolean) GetRef() swagger.Definition {
	return b
}

func Boolean() *boolean {
	return &boolean{}
}


