package api

import (
	"github.com/gin-gonic/gin"
	db "study_lab/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/quizzes", server.createQuiz)
	router.GET("/quizzes/:id", server.getQuiz)
	router.PATCH("/quizzes/:id", server.updateQuiz)
	router.DELETE("/quizzes/:id", server.deleteQuiz)

	router.POST("/quizzes/:id/tags", server.addTagToQuiz)
	router.DELETE("/quizzes/:id/tags", server.removeTagFromQuiz)

	router.POST("/quizzes/:id/questions", server.addQuestion)
	router.GET("/quizzes/:id/questions", server.getQuestions)
	router.PATCH("/questions/:id", server.updateQuestion)
	router.DELETE("/questions/:id", server.deleteQuestion)

	router.POST("/questions/:id/answers", server.addAnswer)
	router.GET("/questions/:id/answers", server.getAnswers)
	router.PATCH("/answers/:id", server.updateAnswer)
	router.DELETE("/answers/:id", server.deleteAnswer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
