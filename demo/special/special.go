package main

type MyType string

func (t MyType) Get(key string) string {
	return "hello"
}

// ServerFunc is a Server func.
type ServerFunc func(params Params) (*Image, error)

// Get implements Server.
func (f ServerFunc) Get(params Params) (*Image, error) {
	return f(params)
}

func main() {
	st := MyType("tag")
	st.Get("key") // <- unresolved Get
}
