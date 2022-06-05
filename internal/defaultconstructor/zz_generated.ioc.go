//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package defaultconstructor

import (
	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/singleton"
)

func init() {
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "AppAlias",
		Factory: func() interface{} {
			return &Hello{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			return New(), nil
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "AppAlias",
		Factory: func() interface{} {
			return &World{}
		},
		ParamFactory: func() interface{} {
			var _ paramInterface = &Param{}
			return &Param{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(paramInterface)
			impl := i.(*World)
			return param.Init(impl)
		},
	})
}

type paramInterface interface {
	Init(impl *World) (*World, error)
}