package handlers

import (
	"URLProject/configs"
	"URLProject/internal/delivery/payload"
	"URLProject/pkg/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthServer struct {
	*configs.Config
}

type AuthServerDeps struct {
	*configs.Config
}

func NewAuthServer(deps AuthServerDeps) *AuthServer {
	return &AuthServer{Config: deps.Config}
}

func (as *AuthServer) RegisterUser(ctx *gin.Context) {
	requestStruct, err := request.HandleBody[payload.RegisterRequest](ctx)
	if err != nil {
		log.Printf("unable to handle request body: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(requestStruct)
}

func (as *AuthServer) LoginUser(ctx *gin.Context) {
	requestStruct, err := request.HandleBody[payload.LoginRequest](ctx)
	if err != nil {
		log.Printf("unable to handle request body: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(requestStruct)

	resp := payload.LoginResponse{
		Token: "098239d796sd862169y9&",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}
