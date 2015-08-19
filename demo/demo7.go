package demo

import "fmt"

type stru struct {
	Field  string
	Field1 map[string]string
}

type str struct {
	Fiel *stru
}

func de() *stru {
	a := &stru{
		Field:  "demo",
		Field1: map[string]string{"f1": "val"},
	}

	a.Field = "dem"
	a.Field1 = map[string]string{"f2": "val"}

	return
}

func dem() {
	b := de()
	b.Field = "de"
	b.Field1 = map[string]string{"f3": "fal"}

	a := *str{
		Fiel: &stru{
			Field:  fmt.Sprint("demo%d", 1),
			Field1: map[string]string{"a", "n"},
		},
	}

	for i := 0; i < 4; i++ {
		a.Fiel = &stru{
			Field:  fmt.Sprintf("demo %d", i),
			Field1: map[string]string{"a", "n"},
		}
	}

	_ = b
	_ = a
}
