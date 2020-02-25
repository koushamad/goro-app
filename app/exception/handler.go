package exception

import (
	"github.com/koushamad/goro/app/exception"
	"net/http"
)

func Handler(e exception.Exception) (int ,interface{}) {
	switch e.Code {
	case 1: // sample
		return http.StatusOK, nil
	default:
		return http.StatusInternalServerError, nil
	}
}
