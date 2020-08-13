package test

import (
	"github.com/goolanger/swaggerize/models/document"
	"github.com/goolanger/swaggerize/models/document/licenses"
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/path"
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/scheme"
	"github.com/goolanger/swaggerize/pkg/io"
	"testing"
)

func TestSwaggerInit(t *testing.T) {
	api := swagger.New().
		Info(document.Info{
			License: licenses.Apache2,
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

	tag := api.Define(model.Object("Tag").Props(
		model.Property("id", model.Int()),
		model.Property("name", model.String()),
	))

	category := api.Define(
		model.Object("Category").Props(
			model.Property("id", model.Int()).
				Tag("x-go-custom-tag", "gorm:\"primary_key;auto_increment:false\""),
			model.Property("lang", model.String()).
				Tag("x-go-custom-tag", "gorm:\"primary_key;auto_increment:false;index:lang\""),
			model.Property("image", model.File()),
			model.Property("name", model.String()),
		),
	)

	offer := api.Define(
		model.Object("Offer").Props(
			model.Property("id", model.Int()).
				Tag("x-go-custom-tag", "gorm:\"primary_key;auto_increment:false\""),
			model.Property("categoryId", model.Int()),
			model.Property("shopId", model.Int()),
			model.Property("name", model.String()),
			model.Property("availability", model.Enum(model.String(), "available", "pending", "sold out")),

			model.Property("quantity", model.Int()).
				Description("how many units available for this offer"),
			model.Property("is_discount", model.Boolean()).
				Description("offers a discount if true, an absolute price if false"),
			model.Property("value", model.Number()).
				Description("either carries a discount percentage or an absolute price depending on is_discount"),

			model.Property("valid_from", model.DateTime()).
				Description("describes the range start of date-time when the offer is available").
				Tag("x-go-custom-tag", "gorm:\"Type:timestamp\""),
			model.Property("valid_until", model.DateTime()).
				Description("describes the range end of date-time when the offer is available").
				Tag("x-go-custom-tag", "gorm:\"Type:timestamp\""),

			model.Property("range", model.Enum(model.String(), "near", "area", "city", "state", "country", "global")).
				Description("the range within this offer will be seen by buyer"),

			model.Property("tags", model.Array(tag.GetRef())),
		),
	)

	shop := api.Define(model.Object("Shop").Props(
		model.Property("id", model.Int()).
			Tag("x-go-custom-tag", "gorm:\"primary_key\""),
		model.Property("name", model.String()).
			Tag("x-go-custom-tag", "gorm:\"unique_index\""),

		model.Property("city", model.String()),
		model.Property("zip", model.String()),
		model.Property("address1", model.String()),
		model.Property("address2", model.String()),
		model.Property("vatId", model.String()),
		model.Property("lat", model.Number()),
		model.Property("lng", model.Number()),
		model.Property("openFrom", model.String()),
		model.Property("openTo", model.String()),

	))

	api.Define(model.Object("Claims").Props(
		model.Property("id", model.Int()),
		model.Property("roles", model.Array(model.String())),
	))

	api.Route(path.Resource(api, category, path.Scope(path.Inherit, path.Inherit).Routes(
		api.Route(path.Resource(api, offer)),
	)))
	api.Route(path.Resource(api, shop))

	err := io.Save(api, "swagger.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
