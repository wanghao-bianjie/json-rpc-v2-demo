package app

import (
	"github.com/osamingo/jsonrpc/v2"
	"log"
	"myjsonrpcv2/internal/app/api"
	"myjsonrpcv2/internal/app/conf"
	"myjsonrpcv2/internal/app/db"
	"myjsonrpcv2/internal/app/global"
	"net/http"
)

func Serve(cfg *conf.Config) {
	global.Config = cfg
	db.InitMysqlDB(cfg.Mysql)
	db.CreateTable()
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
