package main

import (
	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/config"
	"github.com/fatih/color"
	_ "github.com/photowey/helloioc/internal/email"
	"github.com/photowey/helloioc/internal/order"
	_ "github.com/photowey/helloioc/internal/sms"
)

func main() {
	if err := loadIoC(); err != nil {
		panic(err)
	}
	printConfig()

	order.Run()
}

func printConfig() {
	helloioc := ""
	_ = config.LoadConfigByPrefix("hello.ioc", &helloioc)
	color.Blue("the value of configs hello.ioc:%s", helloioc)
}

func loadIoC() error {
	nameOpt := config.WithConfigName("ioc_golang")
	typeOpt := config.WithConfigType("yaml")
	err := config.Load(nameOpt, typeOpt)
	if err != nil {
		panic(err)
	}

	return autowire.Load()
}
