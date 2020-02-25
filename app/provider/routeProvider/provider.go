package routeProvider

import (
	"github.com/koushamad/goro-app/app/http/route"
	"github.com/koushamad/goro-core/app/http/router"
)

func Routes(Router *router.Router) {
	route.Api(Router)
}
