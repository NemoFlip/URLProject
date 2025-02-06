package handlers

import (
	"URLProject/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type StatServer struct {
	statRepository *repository.StatRepository
}

func NewStatServer(statRepository *repository.StatRepository) *StatServer {
	return &StatServer{statRepository: statRepository}
}

// @Summary Statistics
// @Description Get statistics
// @Tags Stat
// @Produce json
// @Param from query string true "From date"
// @Param to query string true "To date"
// @Param by query string true "month or day"
// @Success 200 {object} map[string]string "all stats was found"
// @Failure 400 {object} string "bad credentials"
// @Failure 500 {object} string "internal server error"
// @Router /stat [get]
func (ss *StatServer) GetStatistics(ctx *gin.Context) {
	from, err := time.Parse("2006-01-02", ctx.Query("from"))
	if err != nil {
		log.Println("Invalid from query param")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	to, err := time.Parse("2006-01-02", ctx.Query("to"))
	if err != nil {
		log.Println("Invalid to query param")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	by := ctx.Query("by")
	if by != "day" && by != "month" {
		log.Println("Invalid by query param")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	stats := ss.statRepository.GetStats(by, from, to)

	ctx.JSON(http.StatusOK, stats)
}
