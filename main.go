package main

import (
	"a2sv_stocet_learning_path/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
