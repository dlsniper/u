package demo2094

type demo struct {
	a string
	b int32
	c string
}

type bemo struct {
	demo
	x string
}

func _() (struct {
	x string
}) {
	_ = struct{ x string }{x: "demo"}
	_ = <weak_warning descr="Structure has unnamed field initialization">demo{"demo"}</weak_warning>
b, _ := <weak_warning descr="Structure has unnamed field initialization">demo{"demo"}</weak_warning>, 1
_ = <weak_warning descr="Structure has unnamed field initialization">demo{"demo", 1, "demo"}</weak_warning>
_ = demo{a: "demo"}
_ = demo{a: "demo", b: 1}
_ = <weak_warning descr="Structure has unnamed field initialization">demo{
a: "demo",
1,
}</weak_warning>
_ = bemo{x: "demo"}
_ = b
return <weak_warning descr="Structure has unnamed field initialization">struct{x string}{"demo"}</weak_warning>
}