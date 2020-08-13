package test

import (
	"github.com/goolanger/swaggerize/models/document"
	"github.com/goolanger/swaggerize/models/document/licenses"
	"github.com/goolanger/swaggerize/models/model"
	"github.com/goolanger/swaggerize/models/model/gorm"
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
		gorm.Id(),
		gorm.String("name", ""),
	))

	category := api.Define(
		model.Object("Category").Props(
			gorm.Id(),
			gorm.Index("lang", "", model.String()),
			gorm.File("image",""),
			gorm.String("name", ""),
		),
	)

	shop := api.Define(model.Object("Shop").Props(
		gorm.Id(),
		gorm.Index("name", "", model.String()),

		gorm.String("city",""),
		gorm.String("zip", ""),
		gorm.String("address1", ""),
		gorm.String("address2", ""),
		gorm.String("vatId", ""),
		gorm.Number("lat", ""),
		gorm.Number("lng", ""),
		gorm.DateTime("openFrom", ""),
		gorm.DateTime("openTo", ""),
	))

	offer := api.Define(
		model.Object("Offer").Props(
			gorm.Id(),
			gorm.Reference(category),
			gorm.Reference(shop),

			gorm.String("name", ""),
			gorm.Enum("availability", "",
				model.String(),
				"available",
				"pending",
				"sold out",
			),

			gorm.Int("quantity", "how many units available for this offer"),
			gorm.Boolean("is_discount", "offers a discount if true, an absolute price if false"),
			gorm.Number("value", "either carries a discount percentage or an absolute price depending on is_discount"),
			gorm.DateTime("valid_from", "describes the range start of date-time when the offer is available"),
			gorm.DateTime("valid_until", "describes the range ends of date-time when the offer is available"),

			gorm.Enum("range", "the range within this offer will be seen by buyer",
				model.String(),
				"near",
				"area",
				"city",
				"state",
				"country",
				"global",
			),

			gorm.Array("tags", "", tag.GetRef()),
		),
	)

	api.Define(model.Object("Claims").Props(
		gorm.Id(),
		gorm.Array("roles", "", model.String()),
	))

	api.Route(path.Resource(api, category, path.Scope(path.Inherit, path.Inherit).Routes(
		api.Route(path.Resource(api, offer)),
	)))
	//api.Route(path.Resource(api, shop))

	err := io.Save(api, "swagger.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
