package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	maxSize int

	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, maxSize int, uc Usecase) *Server {
	api := Server{
		maxSize: maxSize,
		uc:      uc,
	}

	api.server = echo.New()
	api.address = fmt.Sprintf("%s:%d", ip, port)

	api.server.Use(middleware.CORS())

	api.server.GET("/quizzes", api.GetAllQuizzes)
	api.server.GET("/quizzes/:id", api.GetQuizByID)
	api.server.GET("/quizzes/:id/questions", api.GetQuestionsByQuizID)
	api.server.POST("/quizzes/:quiz_id/questions", api.CreateQuestion)
	api.server.POST("/quizzes", api.CreateQuiz)

	api.server.POST("/questions/:question_id/answer", api.CreateAnswer)
	api.server.GET("/questions/:question_id/answer", api.GetAnswersByQuestionID)

	api.server.POST("/quizzes/:id/rating", api.SubmitQuizRating)
	api.server.GET("/quizzes/:id/average-rating", api.GetAverageRating)
	api.server.GET("/quizzes/:quiz_id/comments", api.GetCommentsByQuizID)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
