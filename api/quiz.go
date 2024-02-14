package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "study_lab/db/sqlc"
	"time"
)

type createQuizRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createQuiz(ctx *gin.Context) {
	var req createQuizRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	name := req.Name

	quiz, err := server.store.CreateQuiz(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, quiz)
}

type getQuizRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getQuizResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	TagIds    []int64   `json:"tag_ids"`
}

func (server *Server) getQuiz(ctx *gin.Context) {
	var req getQuizRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := req.ID

	quiz, err := server.store.GetQuiz(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tags, err := server.store.GetTagsForQuiz(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getQuizResponse{
		ID:        quiz.ID,
		Name:      quiz.Name,
		CreatedAt: quiz.CreatedAt,
		TagIds:    extractTagIds(tags),
	}

	ctx.JSON(http.StatusOK, res)
}

func extractTagIds(tags []db.Tag) []int64 {
	ids := make([]int64, 0, len(tags))
	for _, tag := range tags {
		ids = append(ids, tag.ID)
	}
	return ids
}
