package main

import (
	"dousheng/router"
	"fmt"
)

func main() {
	r := router.InitRouter()
	err := r.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		return
	}
}
