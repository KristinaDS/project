package usecase

import (
	"backend/pkg/models"
)

type Provider interface {
	CreateQuiz(quiz models.Quiz) (int, error)

	UpdateQuiz(id int, quiz models.Quiz) error

	DeleteQuiz(id int) error

	// GetAllQuizzes возвращает все викторины
	GetAllQuizzes() ([]models.Quiz, error)

	// GetQuizByID возвращает викторину по ID
	GetQuizByID(id int) (models.Quiz, error)

	// Работа с вопросами
	CreateQuestion(id int, questionText string) error
	GetQuestionsByQuizID(id int) ([]models.Questions, error)
}
