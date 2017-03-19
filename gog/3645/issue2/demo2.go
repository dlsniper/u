package issue2

import "github.com/dlsniper/u/gog/3645"

func _() {
	a := issue.Demo{}
	a.Method() // refactor here
	_ = a
}

