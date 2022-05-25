package order

import (
	"reflect"
	"strings"

	"github.com/alibaba/ioc-golang/autowire/singleton"
	"github.com/fatih/color"
	"github.com/photowey/helloioc/internal/sdid"
	"github.com/photowey/helloioc/internal/sms"
)

func Run() {
	CreateOrder()
}

func CreateOrder() bool {

	// isms, err := singleton.GetImpl("Message-Sender")
	sender, err := singleton.GetImpl(sdid.MessageSms) // instead of the hardcode Message-Sender
	if err != nil {
		panic(err)
	}
	smsSender := sender.(*sms.Sender)
	channel := smsSender.Channel()
	color.Blue("message channel is: %s", channel)
	typeOf := reflect.TypeOf(*smsSender)
	pkgPath := typeOf.PkgPath()
	fullName := strings.ReplaceAll(pkgPath, "/", ".")
	color.Blue("type.pkg.path.name is: %s", fullName)

	return true
}
