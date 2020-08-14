package tags

type tag struct {
	Name, Description string
}

func (t *tag) GetRep() map[string]interface{} {
	return map[string]interface{}{
		"description":t.Description,
		"name":t.Name,
	}
}

func (t *tag) GetName() string {
	return t.Name
}

func New(name, description string) *tag {
	return &tag{Name: name, Description: description}
}

