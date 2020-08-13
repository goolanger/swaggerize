package tag

type Tag struct {
	name string
}

func (t Tag) GetName() string {
	return t.name
}

func New(name string) *Tag {
	return &Tag{name: name}
}

