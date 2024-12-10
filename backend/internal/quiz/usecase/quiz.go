package usecase

import (
	"backend/pkg/models"
)

// GetAllQuizzes возвращает все викторины
func (u *Usecase) GetAllQuizzes() ([]models.Quiz, error) {
	return u.p.GetAllQuizzes()
}

// GetQuizByID возвращает викторину по ID
func (u *Usecase) GetQuizByID(id int) (models.Quiz, error) {
	return u.p.GetQuizByID(id)
}

func (u *Usecase) CreateQuiz(quiz models.Quiz) (int, error) {
	return u.p.CreateQuiz(quiz)
}

func (u *Usecase) UpdateQuiz(id int, quiz models.Quiz) error {
	return u.p.UpdateQuiz(id, quiz)
}

func (u *Usecase) DeleteQuiz(id int) error {
	return u.p.DeleteQuiz(id)
}

// Добавление нового вопроса
func (u *Usecase) CreateQuestion(quizID int, questionText string) error {
	return u.p.CreateQuestion(quizID, questionText)
}

// Получение всех вопросов для викторины
func (u *Usecase) GetQuestionsByQuizID(quizID int) ([]models.Questions, error) {
	return u.p.GetQuestionsByQuizID(quizID)
}
