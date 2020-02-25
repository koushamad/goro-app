package controller

import (
	"github.com/koushamad/goro-app/app/http/contrroller/resource"
	"github.com/koushamad/goro-app/app/service"
	"github.com/koushamad/goro-core/app/http/request"
	"github.com/koushamad/goro-core/app/http/response"
	"net/http"
)

func Hello(req *request.Request, res *response.Response) {
	var helloResponse resource.HelloRequest
	req.Pars(&helloResponse)

	name := service.Hello(helloResponse)

	res.Send(http.StatusOK, resource.Success(resource.HelloResponse{
		Name: name,
	}))
}
