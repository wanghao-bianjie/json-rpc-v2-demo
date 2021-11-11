package main

import (
	"log"
	"myjsonrpcv2/internal/app/api"
	_ "myjsonrpcv2/internal/app/api/methods"
	"net/http"

	"github.com/osamingo/jsonrpc/v2"
)

func main() {
	mr := jsonrpc.NewMethodRepository()

	for _, s := range api.GetMethods() {
		for _, h := range s.Handlers() {
			_ = mr.RegisterMethod(s.MethodName(h), h, h.Params(), h.Result())
			log.Printf("registe method:\t%s", s.MethodName(h))
		}
	}

	http.Handle("/jrpc", mr)
	http.HandleFunc("/jrpc/debug", mr.ServeDebug)

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatalln(err)
	}
}
