package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "study_lab/db/sqlc"
)

type addQuestionRequestURI struct {
	QuizID int64 `uri:"id" binding:"required"`
}

type addQuestionRequestJSON struct {
	Title string `json:"title" binding:"required"`
}

func (server *Server) addQuestion(ctx *gin.Context) {
	var reqURI addQuestionRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqJSON addQuestionRequestJSON
	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	question, err := server.store.CreateQuestion(ctx, db.CreateQuestionParams{
		QuizID: reqURI.QuizID,
		Title:  reqJSON.Title,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}

type getQuestionsRequest struct {
	QuizID int64 `uri:"id" binding:"required"`
}

type getQuestionsResponse struct {
	Questions []db.Question `json:"questions"`
}

func (server *Server) getQuestions(ctx *gin.Context) {
	var req getQuestionsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	questions, err := server.store.ListAllQuestions(ctx, req.QuizID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getQuestionsResponse{Questions: questions}
	ctx.JSON(http.StatusOK, res)
}

type updateQuestionRequestURI struct {
	ID int64 `uri:"id" binding:"required"`
}

type updateQuestionRequestJSON struct {
	Title string `json:"title" binding:"required"`
}

func (server *Server) updateQuestion(ctx *gin.Context) {
	var reqURI updateQuestionRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqJSON updateQuestionRequestJSON
	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	question, err := server.store.UpdateQuestion(ctx, db.UpdateQuestionParams{
		ID:    reqURI.ID,
		Title: reqJSON.Title,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}

type deleteQuestionRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteQuestion(ctx *gin.Context) {
	var req deleteQuestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteQuestion(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
