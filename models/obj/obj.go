package obj

import "github.com/goolanger/swaggerize/models/swagger"

type object struct {
	name string
	props []swagger.Definition
}

func (o *object) GetName() string {
	return o.name
}

func (o *object) GetRep() map[string]interface{} {
	rep := map[string]interface{}{
		"type": "object",
	}

	if len(o.props) >0 {
		props := make(map[string]interface{})
		for _, p := range o.props {
			props[p.GetName()] = p.GetRep()
		}
		rep["properties"] = props
	}

	return rep
}

func (o *object) GetRef() swagger.Definition {
	return Reference(o.name, "#/definitions/" + o.name)
}

func Object(name string) *object {
	return &object{name: name}
}

func (o *object) Props(props ...swagger.Definition) *object {
	o.props = append(o.props, props...)
	return o
}

