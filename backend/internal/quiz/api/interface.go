package api

import (
	"backend/pkg/models"
)

type Usecase interface {
	GetAllQuizzes() ([]models.Quiz, error)
	GetQuizByID(id int) (models.Quiz, error)
	CreateQuiz(quiz models.Quiz) (int, error)

	CreateQuestion(quizID int, questionText string) (int, error)
	GetQuestionsByQuizID(quizID int) ([]models.Questions, error)

	CreateAnswer(questionID int, answerText string, isCorrect bool) (int, error)

	GetAnswersByQuestionID(questionID int) ([]models.Answers, error)

	SaveQuizRating(quizID int, rating models.Rating) error
	GetAverageRating(quizID int) (float64, error)
	GetCommentsByQuizID(quizID int) ([]models.Rating, error)
}
