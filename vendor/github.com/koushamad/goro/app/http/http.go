package http

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro/app/exception/throw"
	"sync"
)

type(
	HTTP struct {
		Throw *throw.Throw
	}
)

var (
	Engine *gin.Engine
	once sync.Once
	self *HTTP
)

func Boot(engine *gin.Engine)  {
	once.Do(func() {
		Engine = engine
	})
	Init()
}

func Init()  {
	self = &HTTP{
		throw.Load(),
	}
}

func Load() *HTTP {
	return self
}

func (http *HTTP)Serve(port string)  {
	throw.Fatal("Cannot serve Goro", 112, Engine.Run(port))
}