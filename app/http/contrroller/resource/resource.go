package resource

import "github.com/koushamad/goro/app/exception"

type (
	Msg struct {
		Code  int    `json:"code"`
		Title string `json:"title"`
		Text  string `json:"text"`
	}

	Err struct {
		Code    int    `json:"code"`
		Message string `json:"title"`
		Error   string `json:"text"`
	}

	Response struct {
		Status   bool        `json:"status"`
		Code     int         `json:"code"`
		Messages []Msg       `json:"message"`
		Errors   []Err       `json:"errors"`
		Data     interface{} `json:"data"`
	}
)

func Success(data interface{}) interface{} {
	return Response{
		Status:   true,
		Code:     1,
		Messages: []Msg{},
		Errors:   nil,
		Data:     data,
	}
}

func Error(e exception.Exception) interface{} {
	return Response{
		Status:   false,
		Code:     e.Code,
		Messages: []Msg{},
		Errors:   []Err{{Code: e.Code, Message: e.Message, Error: getError(e)}},
		Data:     nil,
	}
}

func getError(e exception.Exception) string {
	if e.Err == nil {
		return e.Message
	}
	return e.Err.Error()
}
