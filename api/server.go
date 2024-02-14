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

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
