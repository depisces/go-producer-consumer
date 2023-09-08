package main

import (
	many_many "go-producer-consumer/many-many"
	"go-producer-consumer/out"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//获取对象
	o := out.NewOut()

	//开启协程
	go o.OutPut()

	//one_one.Exec()
	//one_many.Exec()
	//many_one.Exec()
	many_many.Exec()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
