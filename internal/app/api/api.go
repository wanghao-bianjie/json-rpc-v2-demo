package api

import (
	export2 "myjsonrpcv2/internal/app/api/export"
)

var methods []export2.Method

func RegisterMethod(method export2.Method) {
	methods = append(methods, method)
}

func GetMethods() []export2.Method {
	return methods
}