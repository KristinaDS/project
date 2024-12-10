package api

import (
	"backend/pkg/models"
)

type Usecase interface {
	GetAllQuizzes() ([]models.Quiz, error)
	GetQuizByID(id int) (models.Quiz, error)
	CreateQuiz(quiz models.Quiz) (int, error)
	UpdateQuiz(id int, quiz models.Quiz) error
	DeleteQuiz(id int) error
}
