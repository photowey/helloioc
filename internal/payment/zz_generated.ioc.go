//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package payment

import (
	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/singleton"
)

func init() {
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: &Payment{},
		Factory: func() interface{} {
			return &Payment{}
		},
	})
}
