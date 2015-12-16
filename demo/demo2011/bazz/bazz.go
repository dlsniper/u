package bazz
import "github.com/dlsniper/u/demo/demo2011/foo"

func f(foo *foo.Foo) int {
	return foo.Bar // variable foo is resolved to package foo resulting in Bar being unresolved
}

