package controllers

import (
	"log"
	"net/http"
	"strconv"

	"majoo/dto"
	"majoo/response"
	"majoo/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
	userService service.UserService
}

func NewAuthHandler(
	authService service.AuthService,
	jwtService service.JWTService,
	userService service.UserService,
) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

// @Summary Login
// @Description Login
// @ID Login
// @Param body body dto.LoginRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/auth/login [post]
func (c *authHandler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBind(&loginRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), response.EmptyObj{})
		log.Println("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if loginRequest.UserName == "" || loginRequest.Password == "" {
		response := response.BuildErrorResponse("Failed to process request", "user_name and password required", response.EmptyObj{})
		log.Println("Failed to process request", "user_name and password required")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = c.authService.VerifyCredential(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		response := response.BuildErrorResponse("Failed to login", err.Error(), response.EmptyObj{})
		log.Println("Failed to login", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, _ := c.userService.FindUserByUserName(loginRequest.UserName)

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token
	response := response.BuildResponse(true, "OK!", user)
	log.Println("login success", user)
	ctx.JSON(http.StatusOK, response)

}

// @Summary Register
// @Description Register
// @ID Register
// @Param body body dto.RegisterRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/auth/register [post]
func (c *authHandler) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest

	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), response.EmptyObj{})
		log.Println("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if registerRequest.UserName == "" || registerRequest.Password == "" {
		response := response.BuildErrorResponse("Failed to register request", "user_name and password required", response.EmptyObj{})
		log.Println("Failed to register request", "user_name and password required")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.CreateUser(registerRequest)
	if err != nil {
		response := response.BuildErrorResponse(err.Error(), err.Error(), response.EmptyObj{})
		log.Println("Failed to login", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token
	response := response.BuildResponse(true, "OK!", user)
	log.Println("register success", user)
	ctx.JSON(http.StatusCreated, response)

}
