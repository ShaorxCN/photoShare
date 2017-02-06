package main

import (
	"github.com/astaxie/beego"
	_ "photoShare/initial"
	_ "photoShare/routers"
)

func main() {
	beego.Run()
}
