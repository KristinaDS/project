package usecase

import (
	"backend/pkg/models"
)

type Provider interface {
	// StartQuiz - сохраняет информацию о начале викторины
	StartQuiz(quizID int, userID int, startedAt string) (int, error)

	// SaveAnswersAndGetScore - сохраняет ответы пользователя и возвращает итоговый балл и время прохождения
	SaveAnswersAndGetScore(participationID int, answers []models.User_answers, completedAt string) (int, int, error)

	GetQuizParticipationByID(id int) (models.QuizParticipation, error)

	// FinishQuiz - обновляет информацию о завершении викторины
	FinishQuiz(participationID int, completedAt string) error

	GetQuizInfo(quizID int) (models.Quiz, error)

	GetAllParticipations(quizID int) ([]models.QuizParticipation, error)
	GetElapsedTime(participationID int) (int, error)
}
