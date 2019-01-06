package main

import (
	"GoModDemo/route"
)

func main() {
	mRoute := route.Route

	mRoute.Run(":8080")
}
