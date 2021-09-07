package infoService

import "github.com/gin-gonic/gin"

func GetInfo(context *gin.Context) (string, int, interface{}) {
	res := gin.H{}
	res["context.Request.URL"] = context.Request.URL
	res["context.DefaultQuery"] = context.DefaultQuery("w1", "0")
	res["context.Query"] = context.Query("w2")
	return "GetInfo", 200, res
}
