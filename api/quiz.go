package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func (server *Server) getQuiz(ctx *gin.Context) {
	// TODO: Implement
}
