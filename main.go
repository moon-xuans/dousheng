package main

import (
	"dousheng/router"
	"fmt"
)

func main() {
	r := router.InitRouter()
	err := r.Run(fmt.Sprintf(":%d", 8888))
	if err != nil {
		return
	}
}
