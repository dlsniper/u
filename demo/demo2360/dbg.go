package main

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	a := 1
	fmt.Print(a)

	beego.Router("/", &MainController{})
	beego.Run()
}
