package test

import (
	"github.com/goolanger/swaggerize/models/document"
	"github.com/goolanger/swaggerize/models/document/licences"
	"github.com/goolanger/swaggerize/models/path"
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/tag"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
	"github.com/goolanger/swaggerize/pkg/io"
	"testing"
)

func TestSwaggerInit(t *testing.T) {
	api := swagger.New().
		Info(document.Info{
		Licence:     licences.Apache2,
		Contact:     document.Contact{
			Name:  "API Support",
			Url:   "http://www.swagger.io/support",
			Email: "support@swagger.io",
		},
		Description: "This is the swagger description",
		Title:       "Swagger title",
		Version:     "0.0.1",
		Terms:       "http://swagger.io/terms/",
	}).
		BasePath("/api").
		ExternalDocs(document.External{
			Description: "For more information read here",
			Url:         "https://swagger.io",
		})

	api.Route(
		path.NewScope("/users").Route(
			api.Route(path.NewRoute(path.Inherit)),
			api.Route(path.NewRoute(path.Inherit).SetMethod(methods.POST)),
			api.Route(
				path.NewScope("/{id}").Route(
					api.Route(path.NewRoute(path.Inherit).SetMethod(methods.PUT)),
					api.Route(path.NewRoute(path.Inherit).SetMethod(methods.DELETE)),
				),
			),
		).
			Produces(mimes.ApplicationJson).
			Consumes(mimes.ApplicationJson).
			Tag(
				api.Tag(tag.New("users")),
				api.Tag(tag.New("scaffold")),
			),
	)
	api.Route(path.NewRoute("/api").Produces(mimes.TextHtml))

	printSwagger(api, t)
}

func TestSwaggerTag(t *testing.T) {

}


func printSwagger(specs *swagger.Instance, t *testing.T) {
	rep, err := io.Encode(specs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Instance Specs:\n", string(rep))
}
