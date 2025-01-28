package handlers

import (
	"URLProject/internal/delivery/payload"
	"URLProject/internal/entity"
	"URLProject/internal/repository"
	"URLProject/pkg/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LinkServer struct {
	linkRepository *repository.LinkRepository
	logger         *log.Logger
}

func NewLinkServer(linkRepository *repository.LinkRepository) *LinkServer {
	return &LinkServer{linkRepository: linkRepository}
}

func (ls *LinkServer) Create(ctx *gin.Context) {
	linkRequest, err := request.HandleBody[payload.LinkCreateRequest](ctx)
	if err != nil {
		log.Printf("unable to get link from request: %s", err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	link := entity.NewLink(linkRequest.Url)

	if err = ls.linkRepository.Create(link); err != nil {
		log.Printf("unable to save link in database: %s", err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, link)
}

func (ls *LinkServer) GoTo(ctx *gin.Context) {

}

func (ls *LinkServer) Update(ctx *gin.Context) {

}

func (ls *LinkServer) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
}
