package http

import "github.com/gin-gonic/gin"

type EmptyBody struct{}

func response[T any](status int, body T, message string) (int, gin.H) {
	response_body := gin.H{
		"status":  status,
		"body":    body,
		"message": message,
	}
	return status, response_body
}

func Ok(ctx *gin.Context, body interface{}, message string) {
	ctx.JSON(response(200, body, message))
}
func Ok_no_body(ctx *gin.Context, message string) {
	ctx.JSON(response(200, EmptyBody{}, message))
}
func Forbidden(ctx *gin.Context, message string) {
	ctx.JSON(response(400, EmptyBody{}, message))
}
func NotFound(ctx *gin.Context, message string) {
	ctx.JSON(response(404, EmptyBody{}, message))
}

func Bad(ctx *gin.Context, message string) {
	ctx.JSON(response(400, EmptyBody{}, message))
}

func Response(ctx *gin.Context, status int, body interface{}, message string) {
	ctx.JSON(response(status, body, message))
}
