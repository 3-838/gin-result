package userService

import "github.com/gin-gonic/gin"

func UpdateUser(context *gin.Context) (string, int, interface{}) {
	res := gin.H{}
	res["context.Request.URL"] = context.Request.URL
	res["context.DefaultQuery"] = context.DefaultQuery("w1", "0")
	res["context.Query"] = context.Query("w2")
	return "UpdateUser", 200, res
}
