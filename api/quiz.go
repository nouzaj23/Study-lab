package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "study_lab/db/sqlc"
	"time"
)

type quizResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []db.Tag  `json:"tags"`
}

type tagRequest struct {
	Name string `json:"name"`
}

type createQuizRequest struct {
	Name string       `json:"name" binding:"required"`
	Tags []tagRequest `json:"tags"`
}

func (server *Server) createQuiz(ctx *gin.Context) {
	var req createQuizRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var quiz db.Quiz
	var tags = make([]db.Tag, len(req.Tags))
	err := server.store.ExecTx(ctx, func(queries *db.Queries) error {
		var err error
		quiz, err = queries.CreateQuiz(ctx, req.Name)
		if err != nil {
			return err
		}

		for _, tagReq := range req.Tags {
			var tag db.Tag
			tag, err = queries.CreateTag(ctx, tagReq.Name)
			if err != nil {
				return err
			}

			err = queries.AddTagToQuiz(ctx, db.AddTagToQuizParams{
				TagID:  tag.ID,
				QuizID: quiz.ID,
			})
			if err != nil {
				return err
			}
			tags = append(tags, tag)
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	res := quizResponse{
		ID:        quiz.ID,
		Name:      quiz.Name,
		CreatedAt: quiz.CreatedAt,
		Tags:      tags,
	}

	ctx.JSON(http.StatusOK, res)
}

type getQuizRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getQuiz(ctx *gin.Context) {
	var req getQuizRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var quiz db.Quiz
	var tags []db.Tag
	err := server.store.ExecTx(ctx, func(queries *db.Queries) error {
		var err error
		quiz, err = server.store.GetQuiz(ctx, req.ID)
		if err != nil {
			return err
		}

		tags, err = server.store.GetTagsForQuiz(ctx, req.ID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := quizResponse{
		ID:        quiz.ID,
		Name:      quiz.Name,
		CreatedAt: quiz.CreatedAt,
		Tags:      tags,
	}

	ctx.JSON(http.StatusOK, res)
}

type updateQuizRequestURI struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateQuizRequestJSON struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateQuiz(ctx *gin.Context) {
	var reqURI updateQuizRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqJSON updateQuizRequestJSON
	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.UpdateQuizParams{
		ID:   reqURI.ID,
		Name: reqJSON.Name,
	}

	var quiz db.Quiz
	var tags []db.Tag
	err := server.store.ExecTx(ctx, func(queries *db.Queries) error {
		var err error
		quiz, err = server.store.UpdateQuiz(ctx, params)
		if err != nil {
			return err
		}

		tags, err = server.store.GetTagsForQuiz(ctx, params.ID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := quizResponse{
		ID:        quiz.ID,
		Name:      quiz.Name,
		CreatedAt: quiz.CreatedAt,
		Tags:      tags,
	}

	ctx.JSON(http.StatusOK, res)
}

type deleteQuizRequestURI struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteQuiz(ctx *gin.Context) {
	var req deleteQuizRequestURI
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := req.ID

	err := server.store.DeleteQuiz(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
