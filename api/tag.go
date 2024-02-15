package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "study_lab/db/sqlc"
)

type addTagToQuizRequestURI struct {
	QuizID int64 `uri:"id" binding:"required"`
}

type addTagToQuizRequestJson struct {
	Tag string `json:"tag"`
}

func (server *Server) addTagToQuiz(ctx *gin.Context) {
	var reqURI addTagToQuizRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqJSON addTagToQuizRequestJson
	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.ExecTx(ctx, func(queries *db.Queries) error {
		tag, err := server.store.CreateOrGetTag(ctx, reqJSON.Tag)
		if err != nil {
			return err
		}

		err = server.store.AddTagToQuiz(ctx, db.AddTagToQuizParams{
			TagID:  tag.ID,
			QuizID: reqURI.QuizID,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

type removeTagFromQuizRequestURI struct {
	QuizID int64 `uri:"id" binding:"required"`
}

type removeTagFromQuizRequestJSON struct {
	Tag string `json:"tag"`
}

func (server *Server) removeTagFromQuiz(ctx *gin.Context) {
	var reqURI removeTagFromQuizRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqJSON removeTagFromQuizRequestJSON
	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.ExecTx(ctx, func(queries *db.Queries) error {
		tag, err := server.store.CreateOrGetTag(ctx, reqJSON.Tag)
		if err != nil {
			return err
		}

		err = server.store.RemoveTagFromQuiz(ctx, db.RemoveTagFromQuizParams{
			TagID:  tag.ID,
			QuizID: reqURI.QuizID,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
