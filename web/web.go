package web

import "github.com/gin-gonic/gin"

type ResultFiner func(ctx *gin.Context) (int, any)

func W(finer ResultFiner) gin.HandlerFunc {
	return func(context *gin.Context) {
		code, result := finer(context)
		if code >= 400 {
			context.JSON(code, gin.H{"error": result.(error).Error()})
		} else if result == nil {
			context.Status(code)
		} else {
			context.JSON(code, result)
		}
	}
}
