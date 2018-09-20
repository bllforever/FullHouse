package main

import (
	_ "FullHouse/routers"
_ "FullHouse/models"
	"github.com/astaxie/beego"
)

func main() {
    beego.Run()
}

