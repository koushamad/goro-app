package event

import (
	"github.com/koushamad/goro-app/app/job"
	"github.com/koushamad/goro-core/app/message"
)

var Listen = [...]string{
	"hello",
}

func Handler(m message.Message) {
	switch m.Header {
	case "hello":
		job.Hello(m)
		break
	}
}
