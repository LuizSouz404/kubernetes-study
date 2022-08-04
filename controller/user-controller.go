package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/luizsouz404/api-rest-go/dto"
	"github.com/luizsouz404/api-rest-go/helper"
	"github.com/luizsouz404/api-rest-go/service"
)

//UserController is a contrat to user actions
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController is creating a new instance of UserController
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (controller *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO

	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	userUpdateDTO.Id = id

	user := controller.userService.Update(userUpdateDTO)

	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}

func (controller *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := controller.jwtService.ValidateToken(authHeader)

	if err != nil {
		panic(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	user := controller.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))

	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}
