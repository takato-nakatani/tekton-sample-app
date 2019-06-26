package main

import (
	"./router"
)

func main()  {
	e := router.Init()
	e.Logger.Fatal(e.Start(":1323"))
}