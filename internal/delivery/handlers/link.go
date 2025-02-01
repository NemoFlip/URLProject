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

// @Summary Create
// @Description Create the link
// @Tags Link
// @Accept json
// @Produce json
// @Param user body payload.LinkCreateRequest true "link's credentials"
// @Success 201 {object} payload.LinkResponse "linke is created successfully"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /link [post]
func (ls *LinkServer) Create(ctx *gin.Context) {
	linkRequest, err := request.HandleBody[payload.LinkCreateRequest](ctx)
	if err != nil {
		log.Printf("unable to get link from request: %s", err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	link := entity.NewLink(linkRequest.Url)
	for {
		existedLink, _ := ls.linkRepository.Get(link.Hash)
		if existedLink == nil {
			break
		}
		link.GenerateHash()
	}

	if err = ls.linkRepository.Create(link); err != nil {
		log.Printf("unable to save link in database: %s", err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, link)
}

// @Summary GoTo
// @Description Get the link by hash
// @Tags Link
// @Accept json
// @Produce json
// @Param hash path string true "URL's Hash"
// @Success 200 {object} payload.LinkResponse "linke was found"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /link/{hash} [get]
func (ls *LinkServer) GoTo(ctx *gin.Context) {
	hash := ctx.Param("hash")
	link, err := ls.linkRepository.Get(hash)
	if err != nil {
		log.Printf("unable to get link by hash: %s", err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(ctx.Writer, ctx.Request, link.Url, http.StatusTemporaryRedirect)
}

// @Summary Update
// @Description Update the link by id
// @Tags Link
// @Accept json
// @Produce json
// @Success 200 {object} payload.LinkResponse "linke is updated successfully"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /link/{id} [patch]
func (ls *LinkServer) Update(ctx *gin.Context) {

}

// @Summary Delete
// @Description Delete the link by id
// @Tags Link
// @Accept json
// @Produce json
// @Success 200 {object} payload.LinkResponse "linke is deleted successfully"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /link/{id} [delete]
func (ls *LinkServer) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
}
