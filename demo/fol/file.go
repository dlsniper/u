package fol

/*
 Demo struc is here, have no
 fear
*/
type Demo struct {
	Field int
}

var (
	Thing = &Demo{Field: 1}
)

func (a *Demo) Dem() {

}

func f1(a, b interface{}) {

}

func F2() {
	defer func() {
		err := recover()
		f1(1, err)
	}()
}
