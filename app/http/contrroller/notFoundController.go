package controller

import (
	"github.com/koushamad/goro-app/app/http/contrroller/resource"
	"github.com/koushamad/goro-core/app/exception"
	"github.com/koushamad/goro-core/app/http/request"
	"github.com/koushamad/goro-core/app/http/response"
	"net/http"
)

func NotFound(req *request.Request, res *response.Response) {
	res.Send(http.StatusNotFound, resource.Error(exception.Init("API not found", 404, nil)))
}
