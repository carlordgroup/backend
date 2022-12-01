package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultFiner func(ctx *gin.Context) (int, any)
type IDResultFiner func(ctx *gin.Context, id int) (int, any)

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

type I struct {
	ID int `uri:"id" binding:"required,number"`
}

func ID(finer IDResultFiner) gin.HandlerFunc {
	return W(func(ctx *gin.Context) (int, any) {
		var i I
		err := ctx.ShouldBindUri(&i)
		if err != nil {
			return http.StatusBadRequest, err
		}
		return finer(ctx, i.ID)
	})
}

type Authenticate interface {
	MustLogin() gin.HandlerFunc
	GetAccount() gin.HandlerFunc
	MustAdmin() gin.HandlerFunc
}
