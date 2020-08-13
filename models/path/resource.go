package path

import (
	"fmt"
	"github.com/goolanger/swaggerize/models/model"
	params "github.com/goolanger/swaggerize/models/parameter"
	"github.com/goolanger/swaggerize/models/response"
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/tags"
	"github.com/goolanger/swaggerize/models/types/locations"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
	"strings"
)

type resource struct {
	scope *scope
	swagger.Path
}

func (res *resource) child(s *scope) swagger.Path {
	res.scope = s
	return s
}

func (res *resource) path(s swagger.Path) *resource {
	res.Path = s
	return res
}

func Resource(api *swagger.Instance, target swagger.Definition, scopes ...*scope) *resource {
	scope := Scope(Inherit, Inherit)
	if len(scopes) > 0 {
		for _, s := range scopes {
			scope.Routes(s.routes...)
			scope.Tag(s.tags...)
			scope.Param(s.params...)
			scope.Produces(s.produces...)
			scope.Consumes(s.consumes...)
			scope.Responds(s.responses...)
		}
	}

	res := &resource{}

	resourceName := strings.ToLower(target.GetName())
	resourceId := strings.ToLower(target.GetName()) + "Id"

	res.path(
		Scope(fmt.Sprintf("/%s", resourceName), fmt.Sprintf(target.GetName())).Routes(
			api.Route(Endpoint(Inherit, "List")).SetMethod(methods.GET).
				Responds(
					response.Response(200, model.Array(target.GetRef())),
				),
			api.Route(Endpoint(Inherit, "Create")).SetMethod(methods.POST).
				Consumes(
					mimes.MultipartFormData,
					mimes.ApplicationJson,
				).
				Param(
					params.Param(resourceName, target.GetRef()).In(locations.BODY),
				).
				Responds(
					response.Response(200, target.GetRef()),
				),
			api.Route(
				res.child(Scope(fmt.Sprintf("/{%s}", resourceId), Inherit).Routes(
					append(
						scope.routes,
						api.Route(Endpoint(Inherit, "Update").SetMethod(methods.PUT)).Param(
							params.Param(resourceName, target.GetRef()).In(locations.BODY),
						),
						api.Route(Endpoint(Inherit, "Destroy").SetMethod(methods.DELETE)),
					)...,
				)),
			).Param(
				params.Int(resourceId).In(locations.PATH),
			).Responds(
				response.Response(200, target.GetRef()),
			),
		).Produces(
			mimes.ApplicationJson,
		).Responds(
			response.Response(500, model.Array(model.String())),
		).Tag(api.Tag(tags.Tag(resourceName, "Default crud resources"))),
	)

	return res
}