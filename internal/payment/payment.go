package payment

import (
	"github.com/alibaba/ioc-golang/autowire/singleton"
	"github.com/photowey/helloioc/internal/order"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Payment struct {
	OrderAliasField    *order.Order `singleton:"order"`
	OrderFullNameField *order.Order `singleton:"github.com/photowey/helloioc/internal/order.Order"`
}

func Run() {
	paymentI, _ := singleton.GetImpl("github.com/photowey/helloioc/internal/payment.Payment")
	payment := paymentI.(*Payment)
	payment.OrderAliasField.SenderSms.Target()
	payment.OrderFullNameField.SenderSms.Target()
}
