package main

import (
	"go-sample/models"
	"go-sample/router"
)

func init() {
	models.Setup()
}

func main() {
	r := router.SetupRouter()
	r.Run() // 在 0.0.0.0:8080 上监听并服务
}
