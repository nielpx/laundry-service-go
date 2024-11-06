
package pkg

import (
    "net/http"
)


type SuccessResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products  interface{}     `json:"data"`
}

type NilResponse struct {
    Code      int             `json:"code"`
	Status    string          `json:"status"`
	Message   string		  `json:"message"`
    Products  interface{}     `json:"data"`
}


func NewResponse(status int, message string, products interface{}) SuccessResponse {
    return SuccessResponse{
        Code: status,
		Status: "Success",
		Message: message,
        Products: products,
    }
}


func NewErrResponse(code int, message string) NilResponse {
    return NilResponse{
		Code:      http.StatusNotFound,
		Status: "error",
		Message: message,
        Products: nil,
    }
}


