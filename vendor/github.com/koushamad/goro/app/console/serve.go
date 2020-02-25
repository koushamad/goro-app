package console

import (
	"github.com/koushamad/goro/app/helper"
	"github.com/koushamad/goro/app/http"
)

func Serve(Command string, args []string)  {
	if Command !=  "" {
		http.Load().Serve(Command)
	}
	http.Load().Serve(helper.Env("PORT"))
}
