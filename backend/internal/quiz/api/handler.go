package api

import (
	"backend/pkg/models"
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

	for i := range quizzes {
		quizzes[i].Questions = nil
	}

	return e.JSON(http.StatusOK, quizzes)
}

// GetQuizByID - Обработчик для получения викторины по ID
func (srv *Server) GetQuizByID(e echo.Context) error {
	idStr := e.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid ID format")
	}

	//fmt.Printf("Fetching quiz with ID: %d\n", id)
	quiz, err := srv.uc.GetQuizByID(id)
	if err != nil {
		//fmt.Printf("Error fetching quiz with ID: %d, Error: %v\n", id, err)
		return e.String(http.StatusInternalServerError, "Error fetching quiz")
	}
	if quiz.ID == 0 {
		return e.String(http.StatusNotFound, "Quiz not found")
	}

	quiz.Questions = nil

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

	return e.JSON(http.StatusCreated, map[string]int{
		"id": quizID,
	})
}

// GetQuestionsByQuizID - Обработчик для получения вопросов по ID викторины
func (srv *Server) GetQuestionsByQuizID(e echo.Context) error {
	idStr := e.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid ID format")
	}

	questions, err := srv.uc.GetQuestionsByQuizID(id)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error fetching questions")
	}

	if len(questions) == 0 {
		return e.String(http.StatusNotFound, "No questions found for this quiz")
	}

	return e.JSON(http.StatusOK, questions)
}

// CreateAnswer - Обработчик для создания нового ответа на вопрос
func (srv *Server) CreateAnswer(e echo.Context) error {

	questionIDStr := e.Param("question_id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid question ID format")
	}

	var answer models.Answers
	if err := e.Bind(&answer); err != nil {
		return e.String(http.StatusBadRequest, "Invalid input data: "+err.Error())
	}

	answerID, err := srv.uc.CreateAnswer(questionID, answer.Answer_text, answer.Is_correct)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error creating answer: "+err.Error())
	}

	return e.JSON(http.StatusCreated, map[string]int{
		"id": answerID,
	})
}

// GetAnswersByQuestionID - Обработчик для получения всех ответов для конкретного вопроса
func (srv *Server) GetAnswersByQuestionID(e echo.Context) error {
	questionIDStr := e.Param("question_id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid question ID format")
	}

	answers, err := srv.uc.GetAnswersByQuestionID(questionID)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error fetching answers: "+err.Error())
	}

	if len(answers) == 0 {
		return e.String(http.StatusNotFound, "No answers found for this question")
	}

	return e.JSON(http.StatusOK, answers)
}

func (srv *Server) CreateQuestion(e echo.Context) error {
	quizIDStr := e.Param("quiz_id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid quiz ID format")
	}

	var question models.Questions
	if err := e.Bind(&question); err != nil {
		return e.String(http.StatusBadRequest, "Invalid input data: "+err.Error())
	}

	questionID, err := srv.uc.CreateQuestion(quizID, question.Question_text)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error creating question: "+err.Error())
	}

	return e.JSON(http.StatusCreated, map[string]int{
		"id": questionID,
	})
}

// SubmitQuizRating - Обработчик для отправки рейтинга викторины
func (srv *Server) SubmitQuizRating(e echo.Context) error {
	quizID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid quiz ID")
	}

	var rating models.Rating
	if err := e.Bind(&rating); err != nil {
		return e.String(http.StatusBadRequest, "Invalid input: "+err.Error())
	}

	err = srv.uc.SaveQuizRating(quizID, rating)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error saving rating: "+err.Error())
	}

	return e.String(http.StatusOK, "Rating submitted successfully")
}

// GetAverageRating - Обработчик для получения среднего рейтинга викторины
func (srv *Server) GetAverageRating(e echo.Context) error {
	quizID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid quiz ID")
	}

	averageRating, err := srv.uc.GetAverageRating(quizID)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error fetching average rating: "+err.Error())
	}

	if averageRating == 0 {
		return e.JSON(http.StatusOK, map[string]interface{}{"average_rating": nil}) // Лучше вернуть null вместо 0
	}

	return e.JSON(http.StatusOK, map[string]float64{"average_rating": averageRating})
}

func (srv *Server) GetCommentsByQuizID(e echo.Context) error {
	quizIDStr := e.Param("quiz_id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid quiz ID")
	}

	comments, err := srv.uc.GetCommentsByQuizID(quizID)
	if err != nil {
		return e.String(http.StatusInternalServerError, "Error fetching comments: "+err.Error())
	}

	if len(comments) == 0 {
		return e.String(http.StatusNotFound, "No ratings found for this quiz")
	}

	return e.JSON(http.StatusOK, comments)
}
