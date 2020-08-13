package test

import (
	"github.com/goolanger/swaggerize/models/document"
	"github.com/goolanger/swaggerize/models/document/licences"
	"github.com/goolanger/swaggerize/models/obj"
	"github.com/goolanger/swaggerize/models/path"
	"github.com/goolanger/swaggerize/models/response"
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
	"github.com/goolanger/swaggerize/models/types/scheme"
	"github.com/goolanger/swaggerize/pkg/io"
	"testing"
)

func TestSwaggerInit(t *testing.T) {
	api := swagger.New().
		Info(document.Info{
			License: licences.Apache2,
			Contact: document.Contact{
				Name:  "API Support",
				Url:   "http://www.swagger.io/support",
				Email: "support@swagger.io",
			},
			Description: "This is the swagger description",
			Title:       "Swagger title",
			Version:     "0.0.1",
			Terms:       "http://swagger.io/terms/",
		}).
		Schemes(scheme.HTTPS, scheme.HTTP).
		Host("localhost:9999").
		BasePath("/api").
		ExternalDocs(document.External{
			Description: "For more information read here",
			Url:         "https://swagger.io",
		})

	tag := api.Define(obj.Object("Tag").Props(
		obj.Property("id", obj.Int()),
		obj.Property("name", obj.String()),
	))

	api.Define(
		obj.Object("Category").Props(
			obj.Property("id", obj.Int()).
				Tag("x-go-custom-tag", "gorm:\"primary_key;auto_increment:false\""),
			obj.Property("lang", obj.String()).
				Tag("x-go-custom-tag", "gorm:\"primary_key;auto_increment:false;index:lang\""),
			obj.Property("image", obj.File()),
			obj.Property("name", obj.String()),
		),
	)

	api.Define(
		obj.Object("Offer").Props(
			obj.Property("id", obj.Int()).
				Tag("x-go-custom-tag", "gorm:\"primary_key;auto_increment:false\""),
			obj.Property("categoryId", obj.Int()),
			obj.Property("shopId", obj.Int()),
			obj.Property("name", obj.String()),
			obj.Property("availability", obj.Enum(obj.String(), "available", "pending", "sold out")),

			obj.Property("quantity", obj.Int()).
				Description("how many units available for this offer"),
			obj.Property("is_discount", obj.Boolean()).
				Description("offers a discount if true, an absolute price if false"),
			obj.Property("value", obj.Number()).
				Description("either carries a discount percentage or an absolute price depending on is_discount"),

			obj.Property("valid_from", obj.DateTime()).
				Description("describes the range start of date-time when the offer is available").
				Tag("x-go-custom-tag", "gorm:\"Type:timestamp\""),
			obj.Property("valid_until", obj.DateTime()).
				Description("describes the range end of date-time when the offer is available").
				Tag("x-go-custom-tag", "gorm:\"Type:timestamp\""),

			obj.Property("range", obj.Enum(obj.String(), "near", "area", "city", "state", "country", "global")).
				Description("the range within this offer will be seen by buyer"),

			obj.Property("tags", obj.Array(tag.GetRef())),
		),
	)

	shop := api.Define(obj.Object("Shop").Props(
		obj.Property("id", obj.Int()).
			Tag("x-go-custom-tag", "gorm:\"primary_key\""),
		obj.Property("name", obj.String()).
			Tag("x-go-custom-tag", "gorm:\"unique_index\""),

		obj.Property("city", obj.String()),
		obj.Property("zip", obj.String()),
		obj.Property("address1", obj.String()),
		obj.Property("address2", obj.String()),
		obj.Property("vatId", obj.String()),
		obj.Property("lat", obj.Number()),
		obj.Property("lng", obj.Number()),
		obj.Property("openFrom", obj.String()),
		obj.Property("openTo", obj.String()),

	))

	api.Define(obj.Object("Claims").Props(
		obj.Property("id", obj.Int()),
		obj.Property("roles", obj.Array(obj.String())),
	))

	api.Route(path.Scope("/shop").Routes(
		api.Route(path.Endpoint(path.Inherit).SetMethod(methods.POST)),
		api.Route(path.Endpoint(path.Inherit).SetMethod(methods.PUT)),
		api.Route(path.Endpoint(path.Inherit).SetMethod(methods.PATCH)),
	).
		Consumes(mimes.ApplicationJson).
		Produces(mimes.ApplicationJson).
		Responds(
			response.Response(200, obj.Array(shop.GetRef())),
			response.Response(404, obj.String()),
			response.Response(500, obj.String()),
		))

	printSwagger(api, t)
}

func printSwagger(specs *swagger.Instance, t *testing.T) {
	rep, err := io.Encode(specs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Instance Specs:\n", string(rep))
}
