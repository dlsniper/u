package ctx

func NewContext() *Context { // this func is marked as not in use
	return &Context{}
}

type (
	Context struct{}
	Context2 struct{}
)

func (c Context) Hello() {
	println("Hello")
}

func (c Context2) Hello() {
	println("Hello2")
}
