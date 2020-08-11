package models

type SwaggerVersion string
const (
	V2 SwaggerVersion = "2.0"
)

type Swagger struct {
	BasePath string
	Definitions map[string] interface{}
	Host string
	Info map[string]interface{}
	Paths map[string] interface{}
	Schemes []string
	Security []interface{}
	SecurityDefinitions map[string] interface{}
	Swagger SwaggerVersion
	Tags []interface{}
}

func New (version SwaggerVersion) *Swagger {
	return &Swagger{Swagger: version}
}

func (s *Swagger) Base(path string) *Swagger {
	s.BasePath = path
	return s
}