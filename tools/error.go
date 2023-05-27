package tools

import (
	"fmt"
	"github.com/pxgo/GoFM/settings"
	"net/http"
)

type IError struct {
	Type   settings.IResponseType
	Status settings.IStatus
	Args   settings.IResponseArgs
}

func (err IError) Error() string {
	return fmt.Sprintf("%d: %s", err.Status, err.Type)
}

type InError struct{}

var Error = InError{}

func (err *InError) newError(status settings.IStatus, resType settings.IResponseType, args settings.IResponseArgs) IError {
	resError := IError{
		Type:   resType,
		Status: status,
		Args:   args,
	}
	return resError
}

func (err *InError) NewError400(resType settings.IResponseType, args settings.IResponseArgs) IError {
	return err.newError(http.StatusBadRequest, resType, args)
}

func (err *InError) NewError500(resType settings.IResponseType, args settings.IResponseArgs) IError {
	return err.newError(http.StatusInternalServerError, resType, args)
}

func (err *InError) NewError403(resType settings.IResponseType, args settings.IResponseArgs) IError {
	return err.newError(http.StatusForbidden, resType, args)
}

func (err *InError) NewCommonError(status settings.IStatus, text string) IError {
	return err.newError(status, settings.ResponseTypes.CommonError, &[]string{text})
}
