package usecase

import (
	"backend/pkg/models"
)

type Provider interface {
	CreateQuiz(quiz models.Quiz) (int, error)

	GetAllQuizzes() ([]models.Quiz, error)

	GetQuizByID(id int) (models.Quiz, error)

	CreateQuestion(quizID int, questionText string) (int, error)
	GetQuestionsByQuizID(id int) ([]models.Questions, error)

	CreateAnswer(questionID int, answerText string, isCorrect bool) (int, error)
	GetAnswersByQuestionID(questionID int) ([]models.Answers, error)

	SaveRating(quizID int, rating models.Rating) error
	GetAverageRating(quizID int) (float64, error)
	GetCommentsByQuizID(quizID int) ([]models.Rating, error)
}
