package configuration

import "github.com/gin-gonic/gin"

type NovelHandlerType func(c *gin.Context) (string, int, interface{})

func NovelHandler() func(n NovelHandlerType) gin.HandlerFunc {
	return func(n NovelHandlerType) gin.HandlerFunc {
		return func(context *gin.Context) {
			msg, code, result := n(context)
			context.JSON(code, gin.H{"message": msg, "code": code, "result": result})
		}
	}
}
