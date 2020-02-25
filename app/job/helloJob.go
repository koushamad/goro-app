package job

import (
	"fmt"
	"github.com/koushamad/goro-cache/app/cache"
	"github.com/koushamad/goro-core/app/message"
)

func Hello(message message.Message) bool {
	var name string
	cache.Get("name", &name)
	fmt.Println(message.GetBody())
	fmt.Println(name)
	return true
}
