package demo1310

func _() {
	func() {} // covered
	"demo"    // covered
	123       // covered
	'c'       // covered
	false     // NOT covered
	1 + 1     // covered incorrectly as it takes the individual values not the expression
	switch 1 {
	}
}

type User interface {
	Name() string
}

func _(users ...User)  { users[0].Name() }