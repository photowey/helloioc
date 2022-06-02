package email

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Sender struct {
}

func (sender *Sender) Channel() string {
	return "email"
}

func (sender *Sender) Content() string {
	return "Hello, world!"
}

func (sender *Sender) Target() string {
	return "photowey@gmail.com"
}
