package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

// Constructor for SayHelloImpl
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

// Constructor for HelloService
func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: sayHello,
	}
}
