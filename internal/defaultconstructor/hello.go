package defaultconstructor

import (
	"fmt"

	"github.com/alibaba/ioc-golang/autowire/singleton"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=New
// +ioc:autowire:alias=HelloAlias

type Hello struct {
}

func (h Hello) SayHello() string {
	return "hello, world"
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=Param.Init
// +ioc:autowire:alias=WorldAlias

type World struct {
	language []string
}

type Param struct {
	Language string
}

func (p *Param) Init(a *World) (*World, error) {
	a.language = append(a.language, p.Language)

	return a, nil
}

func (h World) SayHello() string {
	return "hello, world"
}

func New() *Hello {
	return &Hello{}
}

func Run() {
	helloInf, _ := singleton.GetImpl("HelloAlias")
	hello := helloInf.(*Hello)
	fmt.Println(hello.SayHello())

	worldInf, _ := singleton.GetImpl("WorldAlias")
	world := worldInf.(*World)
	fmt.Println(world.SayHello())
}
