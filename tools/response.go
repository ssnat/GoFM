package tools

import "github.com/pxgo/GoFM/settings"

type IResponse struct {
}

var Response = IResponse{}

func (own *IResponse) GetResponseBody(data interface{}) settings.IResponseBody {
	return settings.IResponseBody{
		Code:    1,
		Type:    "OK",
		Message: "OK",
		Data:    data,
	}
}
