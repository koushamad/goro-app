package middlewar

import "github.com/koushamad/goro/app/http/middleware"

func Api(mid *middleware.Middleware) {
	mid.Set("type", "api")
	mid.Next()
}
