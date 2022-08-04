package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizsouz404/api-rest-go/dto"
	"github.com/luizsouz404/api-rest-go/entity"
	"github.com/luizsouz404/api-rest-go/helper"
	"github.com/luizsouz404/api-rest-go/service"
)

//AuthController is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.ID)
		v.Token = generatedToken
		response := helper.BuildResponse(true, "ok", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Invalid Please check again yout credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createUser := c.authService.CreateUser(dto.UserCreateDTO(registerDTO))
		token := c.jwtService.GenerateToken(createUser.ID)
		createUser.Token = token
		response := helper.BuildResponse(true, "User created", createUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
