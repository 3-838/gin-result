package main

import (
	"github.com/gin-gonic/gin"
	"jiba.cool/gin-result/src/configuration"
	"jiba.cool/gin-result/src/service/infoService"
	"jiba.cool/gin-result/src/service/userService"
)

func main() {
	r := gin.New()
	r.GET("/", configuration.NovelHandler()(infoService.GetInfo))
	r.GET("/GetInfo/:i", configuration.NovelHandler()(infoService.GetInfo))
	r.GET("/UpdateInfo", configuration.NovelHandler()(infoService.UpdateInfo))
	r.GET("/GetUser", configuration.NovelHandler()(userService.GetUser))
	r.GET("/UpdateUser", configuration.NovelHandler()(userService.UpdateUser))
	r.Run(":80")
}
