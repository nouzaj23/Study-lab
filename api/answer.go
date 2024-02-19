package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "study_lab/db/sqlc"
)

type addAnswerRequestURI struct {
	QuestionID int64 `uri:"id" binding:"required"`
}

type addAnswerRequestJSON struct {
	Text      string `json:"text" binding:"required"`
	IsCorrect bool   `json:"is_correct" binding:"required"`
}

func (server *Server) addAnswer(ctx *gin.Context) {
	var reqURI addAnswerRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqJSON addAnswerRequestJSON
	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	answer, err := server.store.CreateAnswer(ctx, db.CreateAnswerParams{
		QuestionID: reqURI.QuestionID,
		Text:       reqJSON.Text,
		IsCorrect:  reqJSON.IsCorrect,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answer)
}

type getAnswersRequestURI struct {
	QuestionID int64 `uri:"id" binding:"required"`
}

type getAnswersResponse struct {
	Answers []db.Answer `json:"answers"`
}

func (server *Server) getAnswers(ctx *gin.Context) {
	var req getAnswersRequestURI
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	answers, err := server.store.ListAnswers(ctx, req.QuestionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getAnswersResponse{
		Answers: answers,
	}
	ctx.JSON(http.StatusOK, res)

}

type updateAnswerRequestURI struct {
	QuestionID int64 `uri:"id" binding:"required"`
}

type updateAnswerRequestJSON struct {
	Text      string `json:"text" binding:"required"`
	IsCorrect bool   `json:"is_correct" binding:"required"`
}

func (server *Server) updateAnswer(ctx *gin.Context) {
	var reqURI addAnswerRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqJSON addAnswerRequestJSON
	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	answer, err := server.store.UpdateAnswer(ctx, db.UpdateAnswerParams{
		ID:        reqURI.QuestionID,
		Text:      reqJSON.Text,
		IsCorrect: reqJSON.IsCorrect,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answer)
}

type deleteAnswerRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteAnswer(ctx *gin.Context) {
	var req deleteAnswerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteAnswer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
