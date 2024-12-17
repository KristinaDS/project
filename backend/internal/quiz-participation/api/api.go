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
	api.server.POST("/quiz_participation/start", api.StartQuizHandler)
	api.server.POST("/quiz_participation/finish", api.FinishQuizHandler)
	api.server.GET("/quiz-participation/:id", api.GetQuizParticipationByID)
	api.server.GET("/quizzes/:id/quiz-participation", api.GetAllParticipations)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
