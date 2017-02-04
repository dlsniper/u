package demo1691

type foo struct{}

var foo int // Should be an error "foo redeclared in this block"
type foo map[int]string  // Should be an error "foo redeclared in this block"
