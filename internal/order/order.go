package order

import (
	"github.com/alibaba/ioc-golang/autowire/singleton"
	"github.com/fatih/color"
	"github.com/photowey/helloioc/internal/email"
	"github.com/photowey/helloioc/internal/message"
	"github.com/photowey/helloioc/internal/sdid"
	"github.com/photowey/helloioc/internal/sms"
)

func Run() {
	CreateOrder()
	sender, _ := singleton.GetImpl("github.com/photowey/helloioc/internal/order.Order")
	order := sender.(*Order)
	order.Send()

}

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:alias=order

type Order struct {
	SenderSms   message.Message `singleton:"github.com/photowey/helloioc/internal/sms.Sender"`
	SenderEmail message.Message `singleton:"github.com/photowey/helloioc/internal/email.Sender"`
}

func CreateOrder() bool {
	handleSms()
	handleEmail()

	return true
}

func handleSms() {
	sender, err := singleton.GetImpl(sdid.MessageSms)
	if err != nil {
		panic(err)
	}
	smsSender := sender.(*sms.Sender)
	target := smsSender.Target()
	color.Blue("message sms channel target is: %s", target)
}

func handleEmail() {
	sender, err := singleton.GetImpl(sdid.MessageEmail)
	if err != nil {
		panic(err)
	}
	emailSender := sender.(*email.Sender)
	target := emailSender.Target()
	color.Blue("message email channel target is: %s", target)
}

func (order Order) Send() bool {
	smsChannelTarget := order.SenderSms.Target()
	emailChannelTarget := order.SenderEmail.Target()
	color.Blue("message sms channel target is: %s", smsChannelTarget)
	color.Blue("message email channel target is: %s", emailChannelTarget)

	return true
}
