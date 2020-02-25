package service

import (
	"github.com/koushamad/goro-app/app/http/contrroller/resource"
	"github.com/koushamad/goro-cache/app/cache"
	"github.com/koushamad/goro-core/app/message"
	"github.com/koushamad/goro-queue/app/queue"
	"time"
)

func Hello(req resource.HelloRequest) string {
	cache.Put("name", req.Name, 60*time.Hour)
	queue.Push(message.Message{Header: "hello", ContentType: message.JSON, Body: []byte(req.Name)})
	return req.Name
}
