package main

import (
	"myjsonrpcv2/cmd"
	_ "myjsonrpcv2/internal/app/api/methods"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 8)
	cmd.Execute()
}
