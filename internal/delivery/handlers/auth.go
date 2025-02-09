package handlers

import (
	"URLProject/configs"
	"URLProject/internal/delivery/payload"
	"URLProject/internal/delivery/services"
	"URLProject/pkg/jwt"
	"URLProject/pkg/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthServer struct {
	*configs.Config
	authService *services.AuthService
}

type AuthServerDeps struct {
	*configs.Config
	*services.AuthService
}

func NewAuthServer(deps AuthServerDeps) *AuthServer {
	return &AuthServer{Config: deps.Config, authService: deps.AuthService}
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
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if _, err = as.authService.Register(requestStruct.Email, requestStruct.Password, requestStruct.Name); err != nil {
		log.Printf("unable to register user: %s", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
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
		log.Printf("unable to handle request body: %s\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	email, err := as.authService.Login(requestStruct.Email, requestStruct.Password)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	jwtStruct := jwt.NewJWT(as.Config.Auth.SecretKey)
	jwtToken, err := jwtStruct.Create(jwt.JWTPayload{Email: email})
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := payload.LoginResponse{Token: jwtToken}
	ctx.JSON(http.StatusOK, resp)
}
