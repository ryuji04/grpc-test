package main

import (
	"grpc-test/controller"
)

func main() {
	//controllerをオブジェクト化
	cont := controller.NewController()
	cont.Execute()
}
