package route

import (
	controller "github.com/koushamad/goro-app/app/http/contrroller"
	"github.com/koushamad/goro-app/app/http/middlewar"
	"github.com/koushamad/goro/app/http/router"
	"github.com/koushamad/goro/app/http/router/group"
)

func Api(Route *router.Router) {
	Route.GroupMiddleware("/api", middlewar.Api, func(Route *group.Router) {
		Route.Post("/hello", controller.Hello)
	})

	Route.Any(controller.NotFound)
}
