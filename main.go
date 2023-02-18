package main

import (
	"dousheng/config"
	"dousheng/router"
	"fmt"
)

func main() {
	r := router.InitRouter()
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port))
	if err != nil {
		return
	}
}
