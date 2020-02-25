package service

import (
	"github.com/koushamad/goro-app/app/http/contrroller/resource"
	"github.com/koushamad/goro-core/app/message"
	"github.com/koushamad/goro-queue/app/queue"
)

func Hello(req resource.HelloRequest) string {
	queue.Push(message.Message{Header: "hello", ContentType: message.JSON, Body: []byte(req.Name)})
	//cache.Put("name", req.Name, 60)
	return req.Name
}
