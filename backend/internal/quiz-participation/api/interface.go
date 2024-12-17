package api

import (
	"backend/pkg/models"
)

type Usecase interface {
	StartQuiz(quizID int, userID int) (int, error)
	FinishQuizU(participationID int, answers []models.User_answers, completedAt string) (int, int, error)
	GetQuizParticipationByID(id int) (models.QuizParticipation, error)
	SaveAnswersAndGetScore(participationID int, answers []models.User_answers, completedAt string) (int, int, error)
	GetAllParticipations(quizID int) ([]models.QuizParticipation, error)
}
