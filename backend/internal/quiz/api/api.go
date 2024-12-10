package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
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

	// Роуты
	api.server.GET("/quizzes", api.GetAllQuizzes)
	api.server.GET("/quizzes/:id", api.GetQuizByID)
	api.server.POST("/quizzes", api.CreateQuiz)
	api.server.PUT("/quizzes/:id", api.UpdateQuiz)
	api.server.DELETE("/quizzes/:id", api.DeleteQuiz)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
