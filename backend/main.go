package main

import (
	"c0ding-backend/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.ApiRouter(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln("服务器启动失败", err)
	}
}
