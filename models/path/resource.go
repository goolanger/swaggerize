package path

import (
	"fmt"
	"github.com/goolanger/swaggerize/models/model"
	params "github.com/goolanger/swaggerize/models/parameter"
	"github.com/goolanger/swaggerize/models/path/restful"
	"github.com/goolanger/swaggerize/models/response"
	"github.com/goolanger/swaggerize/models/swagger"
	"github.com/goolanger/swaggerize/models/types/locations"
	"github.com/goolanger/swaggerize/models/types/methods"
	"github.com/goolanger/swaggerize/models/types/mimes"
	"strings"
)

type resource struct {
	scope *scoped
	swagger.Path
}

func (res *resource) child(s *scoped) swagger.Path {
	res.scope = s
	return s
}

func (res *resource) path(s swagger.Path) *resource {
	res.Path = s
	return res
}

func Resource(api *swagger.Instance, target swagger.Definition, actions *restful.Actions, scopes ...*scoped) *resource {
	scope := Scope(Inherit, Inherit)
	if len(scopes) > 0 {
		for _, s := range scopes {
			scope.Routes(s.routes...)
			scope.Tag(s.tags...)
			scope.Params(s.params...)
			scope.Produces(s.produces...)
			scope.Consumes(s.consumes...)
			scope.Secure(s.secures...)
			scope.Responds(s.responses...)
		}
	}

	res := &resource{}

	resourceName := strings.ToLower(target.GetName()[:1]) + target.GetName()[1:]
	resourceId := resourceName + "Id"

	//resourceTag := api.Tag(tags.New(resourceName, "crud Actions for resource "+resourceName))

	var (
		internalServerError = response.Response(500, "internal server error").
			Schema(model.String())

		notFound = response.Response(404, "not found")

		forbidden = response.Response(403, "forbidden")

		listResourceOk = response.Response(200, "list resources of type "+resourceName).
			Schema(model.Array(target.GetRef()))

		resourceOk = response.Response(200, "fetch resource of type "+resourceName).
			Schema(target.GetRef())

		createResourceOk = response.Response(201, "created "+resourceName).
			Schema(target.GetRef())

		updateResourceOk = response.Response(205, "updated "+resourceName).
			Schema(target.GetRef())

		deleteResourceOk = response.Response(205, "deleted "+resourceName)
	)

	var (
		getRoute    swagger.Path = &scoped{}
		fetchRoute  swagger.Path = &scoped{}
		postRoute   swagger.Path = &scoped{}
		putRoute    swagger.Path = &scoped{}
		deleteRoute swagger.Path = &scoped{}
	)

	if actions.HasGet {
		getRoute = api.Route(Get(Inherit, "Get")).
			Responds(listResourceOk)
	}

	if actions.HasFetch {
		fetchRoute = api.Route(Get(Inherit, "Fetch")).
			Responds(resourceOk)
	}

	if actions.HasPost {
		postRoute = api.Route(Endpoint(Inherit, "Post")).SetMethod(methods.POST).
			Consumes(mimes.MultipartFormData, mimes.ApplicationJson).
			Params(params.Param(resourceName, target.GetRef()).In(locations.BODY)).
			Responds(createResourceOk)
	}

	if actions.HasPut {
		putRoute = api.Route(Endpoint(Inherit, "Put").SetMethod(methods.PUT)).Params(
			params.Param(resourceName, target.GetRef()).In(locations.BODY),
		).Responds(
			updateResourceOk,
		)
	}

	if actions.HasDelete {
		deleteRoute = api.Route(Endpoint(Inherit, "Delete").SetMethod(methods.DELETE)).Responds(
			deleteResourceOk,
		)
	}

	res.path(
		Scope(fmt.Sprintf("/%s", resourceName), fmt.Sprintf(target.GetName())).Routes(
			getRoute, postRoute,
			api.Route(
				res.child(
					Scope(fmt.Sprintf("/{%s}", resourceId), Inherit).Routes(
						append(scope.routes, fetchRoute, putRoute, deleteRoute)...),
				),
			).
				Responds(notFound).
				Params(params.Param(resourceId, model.Long()).In(locations.PATH)),
		).
			Produces(mimes.ApplicationJson).
			Responds(internalServerError, forbidden),
		//Tag(resourceTag),
	)

	return res
}