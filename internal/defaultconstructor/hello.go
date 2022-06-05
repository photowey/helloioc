package defaultconstructor

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=New
// +ioc:autowire:alias=AppAlias

type Hello struct {
}

func (h Hello) SayHello() string {
	return "hello"
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=Param.Init
// +ioc:autowire:alias=AppAlias

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

func (h World) world() string {
	return "world"
}

func New() *Hello {
	return &Hello{}
}
