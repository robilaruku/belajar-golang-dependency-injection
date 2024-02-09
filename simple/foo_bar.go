package simple

type Foo struct {
}

// Constructor for Foo
func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

// Constructor for Bar
func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}
