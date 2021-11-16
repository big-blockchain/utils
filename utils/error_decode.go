package utils

import (
	"google.golang.org/grpc/status"
)

const (
	ErrorCodeNormal  = 200
	UnknownErrorCode = 500
)

type ResponseMap struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrorDecode struct {
}

func (ErrorDecode) ErrorToResponseMap(response interface{}, err error) ResponseMap {
	resp := ResponseMap{}
	if err == nil {
		resp.Code = ErrorCodeNormal
		resp.Msg = "success"
		resp.Data = response
		return resp
	}

	errMap, ok := status.FromError(err)
	if ok == false {
		resp.Code = UnknownErrorCode
		resp.Msg = err.Error()
		resp.Data = response
		return resp
	}

	if errMap == nil {
		resp.Code = ErrorCodeNormal
		resp.Msg = "success"
		resp.Data = response
		return resp
	}
	resp.Code = int(errMap.Code())
	resp.Msg = errMap.Message()
	resp.Data = response

	return resp
}
