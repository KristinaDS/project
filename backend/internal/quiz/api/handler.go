package api

import (
	"backend/pkg/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAllQuizzes - Обработчик для получения всех викторин
func (srv *Server) GetAllQuizzes(e echo.Context) error {
	quizzes, err := srv.uc.GetAllQuizzes()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, quizzes)
}

// GetQuizByID - Обработчик для получения викторины по ID
func (srv *Server) GetQuizByID(e echo.Context) error {
	idStr := e.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Возвращаем ошибку, если ID не является числом
		return e.String(http.StatusBadRequest, "Invalid ID format")
	}

	fmt.Printf("Fetching quiz with ID: %d\n", id)
	quiz, err := srv.uc.GetQuizByID(id)
	if err != nil {
		fmt.Printf("Error fetching quiz with ID: %d, Error: %v\n", id, err)
		return e.String(http.StatusInternalServerError, "Error fetching quiz")
	}

	if quiz.ID == 0 {
		return e.String(http.StatusNotFound, "Quiz not found")
	}
	return e.JSON(http.StatusOK, quiz)
}

// CreateQuiz - Обработчик для создания новой викторины
func (srv *Server) CreateQuiz(e echo.Context) error {
	var quiz models.Quiz
	if err := e.Bind(&quiz); err != nil {
		return e.String(http.StatusBadRequest, "Invalid input data: "+err.Error())
	}
	quizID, err := srv.uc.CreateQuiz(quiz)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error creating quiz: "+err.Error())
	}
	// Возвращаем успешный ответ с ID созданной викторины
	return e.JSON(http.StatusCreated, map[string]int{
		"id": quizID,
	})
}

// UpdateQuiz - Обработчик для обновления викторины по ID
func (srv *Server) UpdateQuiz(e echo.Context) error {
	idStr := e.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid ID format")
	}

	var quiz models.Quiz
	if err := e.Bind(&quiz); err != nil {
		return e.String(http.StatusBadRequest, "Invalid input data: "+err.Error())
	}
	err = srv.uc.UpdateQuiz(id, quiz)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "Quiz updated successfully")
}

// DeleteQuiz - Обработчик для удаления викторины по ID
func (srv *Server) DeleteQuiz(e echo.Context) error {
	idStr := e.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid ID format")
	}

	err = srv.uc.DeleteQuiz(id)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusNoContent, "Quiz deleted successfully")
}
