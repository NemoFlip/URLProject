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

// @Summary Register
// @Description Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body payload.RegisterRequest true "user's info"
// @Success 201 {object} payload.RegisterResponse "user is registered successfully"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /auth/register [post]
func (as *AuthServer) RegisterUser(ctx *gin.Context) {
	requestStruct, err := request.HandleBody[payload.RegisterRequest](ctx)
	if err != nil {
		log.Printf("unable to handle request body: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(requestStruct)
}

// @Summary Login
// @Description Login the user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body payload.LoginRequest true "user's credentials"
// @Success 200 {object} payload.LoginResponse "user is logged in successfully"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /auth/login [post]
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
