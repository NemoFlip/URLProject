package handlers

import (
	"URLProject/internal/delivery/middleware"
	"URLProject/internal/delivery/payload"
	"URLProject/internal/entity"
	"URLProject/internal/repository"
	"URLProject/pkg/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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
		existedLink, _ := ls.linkRepository.GetByHash(link.Hash)
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
	link, err := ls.linkRepository.GetByHash(hash)
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
// @Security BearerAuth
// @Param id path int true "URL's id"
// @Success 200 {object} payload.LinkResponse "linke is updated successfully"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /link/{id} [patch]
func (ls *LinkServer) Update(ctx *gin.Context) {
	linkRequest, err := request.HandleBody[payload.LinkUpdateRequest](ctx)
	if err != nil {
		log.Printf("unable to get struct from body: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if email, ok := ctx.Value(middleware.ContextEmailKey).(string); !ok {
		log.Printf("unable to convert email to string from context: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Println(email)
	}

	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		log.Printf("unable to convert string id to int: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	link := &entity.Link{
		Model: gorm.Model{
			ID: uint(id),
		},
		Url:  linkRequest.Url,
		Hash: linkRequest.Hash,
	}
	if err = ls.linkRepository.Update(link); err != nil {
		log.Printf("unable to update the link by id(%d): %s", id, err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, link)
}

// @Summary Delete
// @Description Delete the link by id
// @Tags Link
// @Accept json
// @Produce json
// @Param id path int true "URL's id"
// @Success 200 {object} payload.LinkResponse "linke is deleted successfully"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /link/{id} [delete]
func (ls *LinkServer) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		log.Printf("unable to convert string id to int: %s", err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = ls.linkRepository.GetByID(uint(id)); err != nil {
		log.Printf("unable to find link in db by id(%d): %s", id, err)
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	if err = ls.linkRepository.Delete(uint(id)); err != nil {
		log.Printf("unable to delete link by id(%d): %s", id, err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}
