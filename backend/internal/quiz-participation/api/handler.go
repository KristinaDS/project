package api

import (
	"backend/pkg/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// StartQuizHandler - Обработчик для старта викторины
func (srv *Server) StartQuizHandler(e echo.Context) error {
	quizID, err := strconv.Atoi(e.QueryParam("quiz_id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid quiz_id")
	}

	userID, err := strconv.Atoi(e.QueryParam("user_id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid user_id")
	}

	participationID, err := srv.uc.StartQuiz(quizID, userID)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error starting quiz: "+err.Error())
	}

	return e.JSON(http.StatusOK, map[string]int{"participation_id": participationID})
}

// FinishQuizHandler - Обработчик для завершения викторины
func (srv *Server) FinishQuizHandler(e echo.Context) error {
	var answers []models.User_answers
	if err := e.Bind(&answers); err != nil {
		return e.String(http.StatusBadRequest, "Invalid input: "+err.Error())
	}

	participationID, err := strconv.Atoi(e.QueryParam("participation_id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid participation_id")
	}

	completedAt := time.Now().Format(time.RFC3339)

	// Сохраняем ответы и получаем баллы и время
	totalScore, elapsedTime, err := srv.uc.FinishQuizU(participationID, answers, completedAt)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error finishing quiz: "+err.Error())
	}

	return e.JSON(http.StatusOK, map[string]int{"total_score": totalScore, "elapsed_time": elapsedTime})
}

// GetQuizParticipationByID - Обработчик для получения информации о прохождении викторины
func (srv *Server) GetQuizParticipationByID(e echo.Context) error {
	// Извлекаем параметр ID из URL
	idStr := e.Param("id")

	// Преобразуем ID в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Если ID невалидный, возвращаем ошибку
		return e.String(http.StatusBadRequest, "Invalid participation ID format")
	}

	// Получаем информацию о прохождении викторины
	quizParticipation, err := srv.uc.GetQuizParticipationByID(id)
	if err != nil {
		// Если возникла ошибка при получении данных, возвращаем ошибку
		return e.String(http.StatusInternalServerError, fmt.Sprintf("Error fetching participation: %v", err))
	}

	// Если данных нет, возвращаем ошибку 404
	if quizParticipation.ID == 0 {
		return e.String(http.StatusNotFound, "Participation not found")
	}

	// Возвращаем информацию о прохождении викторины в формате JSON
	return e.JSON(http.StatusOK, quizParticipation)
}

// GetAllParticipations - Обработчик для получения всех прохождений викторины
func (srv *Server) GetAllParticipations(e echo.Context) error {
	quizID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid quiz ID")
	}

	// Получаем все прохождения викторины
	participations, err := srv.uc.GetAllParticipations(quizID)
	if err != nil {
		return e.String(http.StatusInternalServerError, fmt.Sprintf("Error fetching participations: %v", err))
	}

	// Если данных нет, возвращаем 404
	if len(participations) == 0 {
		return e.String(http.StatusNotFound, "No participations found for this quiz")
	}

	// Возвращаем список всех прохождений в формате JSON
	return e.JSON(http.StatusOK, participations)
}
