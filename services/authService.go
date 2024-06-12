package services

import (
	"reqship-api/helpers/auth"
	"reqship-api/helpers/http"
	"reqship-api/models"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	user := models.LoginUser{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	err := user.Login()

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	token, err := auth.GenerateJWT(user.Username)

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	body := models.UserLoginResponse{
		Token: token,
	}

	http.Ok(ctx, body, "successfully logged in")
}

func SignUp(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.BindJSON(&user); err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	err := user.SignUp()

	if err != nil {
		http.Bad(ctx, err.Error())
		return
	}

	http.Ok_no_body(ctx, "successfully signed up")
}

func CheckAuth(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	username := ctx.Param("username")
	err := auth.CheckJWT(token, &username)

	if err != nil {
		http.Bad(ctx, "invalid token")
		return
	}
	http.Ok_no_body(ctx, "JWT is valid")
}
