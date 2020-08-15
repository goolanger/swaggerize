package test

import (
	"github.com/goolanger/swaggerize/models/document"
	"github.com/goolanger/swaggerize/models/document/licenses"
	"github.com/goolanger/swaggerize/models/model"
	params "github.com/goolanger/swaggerize/models/parameter"
	"github.com/goolanger/swaggerize/models/path"
	"github.com/goolanger/swaggerize/models/response"
	"github.com/goolanger/swaggerize/models/security"
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/tags"
	"github.com/goolanger/swaggerize/models/types/mimes"
	"github.com/goolanger/swaggerize/models/types/scheme"
	"github.com/goolanger/swaggerize/pkg/io"
	"testing"
)

func TestPetStore(t *testing.T) {
	api := swagger.New().
		Info(document.Info{
			License: licenses.Apache2,
			Contact: document.Contact{
				Name:  "goolanger",
				Url:   "https://github.com/goolanger",
				Email: "amauryuh@gmail.com",
			},
			Description: "This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.",
			Title:       "Swagger Petstore",
			Version:     "1.0.5",
			Terms:       "http://swagger.io/terms/",
		}).
		Schemes(scheme.HTTPS, scheme.HTTP).
		Host("petstore.swagger.io").
		BasePath("/v2").
		ExternalDocs(document.External{
			Description: "Find out more about Swagger",
			Url:         "http://swagger.io",
		})

	// Security
	bearer := api.Secure(security.ApiKey("bearer")).GetRef()

	// Tags
	petTag := api.Tag(tags.New("pet", "Everything about your Pets"))
	storeTag :=  api.Tag(tags.New("store", "Access to Petstore orders"))
	userTag :=  api.Tag(tags.New("user", "Operations about user"))

	// Definitions
	apiResponse := api.Define(model.Object("ApiResponse").Props(
		model.Property("code", model.Int()),
		model.Property("type", model.String()),
		model.Property("message", model.String()),
	))

	category := api.Define(model.Object("Category").Props(
		model.Property("id", model.Int64()),
		model.Property("name", model.String()),
	))

	tag := api.Define(model.Object("New").Props(
		model.Property("id", model.Int64()),
		model.Property("name", model.String()),
	))

	pet := api.Define(model.Object("Pet").Props(
		model.Property("id", model.Int64()),
		model.Property("category", category.GetRef()),
		model.Property("name", model.String()).
			Example("doggie"),
		model.Property("photoUrls", model.Array(model.String())),
		model.Property("tags", model.Array(tag.GetRef())),
		model.Property("status", model.Enum(model.String(), "available", "pending", "sold")),
	))

	order := api.Define(model.Object("Order").Props(
		model.Property("id", model.Int64()),
		model.Property("petId", model.Int64()),
		model.Property("quantity", model.Int()),
		model.Property("shipDate", model.DateTime()),
		model.Property("status", model.Enum(model.String(), "placed", "approved", "delivered")),
		model.Property("complete", model.Boolean()),
	))

	//Pets scope
	api.Route(path.Scope("/pet", path.Inherit).Routes(
		api.Route(path.Scope("/{petId}", path.Inherit).Routes(
			api.Route(path.Post("/uploadImage", "uploadFile")).
				Consumes(mimes.MultipartFormData).
				Params(params.Form("additionalMetadata", model.String())).
				Params(params.Form("file", model.File())),

			api.Route(path.Get(path.Inherit, "getPetById")).
				Responds(response.Response(200, "successful operation").Schema(pet.GetRef())).
				Responds(response.Response(400, "invalid id supplied").Schema(model.String())).
				Responds(response.Response(404, "pet not found").Schema(model.String())),

			api.Route(path.Post(path.Inherit, "updatePetWithForm")).
				Params(params.Form("name", model.String())).
				Params(params.Form("status", model.String())).
				Responds(response.Response(200, "successful operation").Schema(pet.GetRef())).
				Responds(response.Response(400, "invalid id supplied").Schema(model.String())).
				Responds(response.Response(404, "pet not found").Schema(model.String())),

			api.Route(path.Delete(path.Inherit, "deletePet")).
				Responds(response.Response(200, "successful operation").Schema(pet.GetRef())).
				Responds(response.Response(400, "invalid id supplied").Schema(model.String())).
				Responds(response.Response(404, "pet not found").Schema(model.String())).
				Secure(security.None()),
		)).
			Produces(mimes.ApplicationJson).
			Params(params.Path("petId", model.Int64())).
			Responds(response.Response(200, "successful operation").Schema(apiResponse.GetRef())),

	)).Tag(petTag).Secure(bearer)

	//Stores scope
	api.Route(path.Scope("/store", "Store").Routes(

	)).Tag(storeTag)

	// Users scope
	api.Route(path.Scope("/user", "User").Routes(

	)).Tag(userTag)

	t.Log(order.GetName())

	err := io.Save(api, "swagger.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
