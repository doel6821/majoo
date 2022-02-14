package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"majoo/response"
	"majoo/service"
	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	TransactionReport(ctx *gin.Context)
	
}

type transactionHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
	userService service.UserService
	transactionService service.TransactionService
}

func NewTransactionHandler(
	authService service.AuthService,
	jwtService service.JWTService,
	userService service.UserService,
	transactionService service.TransactionService,
) TransactionHandler {
	return &transactionHandler{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
		transactionService: transactionService,
	}
}

// @Summary Transaction Report
// @Description Get data transaction 
// @ID Transaction
// @Param dateFrom path string true "datefrom of the transaction to be find"
// @Param dateTo path string true "dateto of the transaction to be find"
// @Param page path string true "page of the transaction to be find"
// @Param limit path string true "limit per page of the transaction to be find"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /transaction/{dateFrom}/{dateTo}/page/limit  [get]
func (c *transactionHandler) TransactionReport(ctx *gin.Context) {
	page := ctx.Param("page")
	limit := ctx.Param("limit")
	p ,_  := strconv.Atoi(page)
	l ,_  := strconv.Atoi(limit)
	layoutFormat := "2006-01-02"
	dateFrom, err := time.Parse(layoutFormat, ctx.Param("dateFrom"))
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), response.EmptyObj{})
		log.Println("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	dateTo, err := time.Parse(layoutFormat, ctx.Param("dateTo"))
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), response.EmptyObj{})
		log.Println("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, _ := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	id ,_  := strconv.Atoi(userID)
	fmt.Println(dateFrom,dateTo,page,limit,id)
	res, err := c.transactionService.TransactionReport(dateFrom, dateTo, int64(id),p,l)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), response.EmptyObj{})
		log.Println("Failed to process request :", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

