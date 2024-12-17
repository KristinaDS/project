package usecase

import (
	"backend/pkg/models"
)

func (u *Usecase) GetAllQuizzes() ([]models.Quiz, error) {
	return u.p.GetAllQuizzes()
}

func (u *Usecase) GetQuizByID(id int) (models.Quiz, error) {
	return u.p.GetQuizByID(id)
}

func (u *Usecase) CreateQuiz(quiz models.Quiz) (int, error) {
	return u.p.CreateQuiz(quiz)
}

func (u *Usecase) CreateQuestion(quizID int, questionText string) (int, error) {
	return u.p.CreateQuestion(quizID, questionText)
}

func (u *Usecase) GetQuestionsByQuizID(quizID int) ([]models.Questions, error) {
	return u.p.GetQuestionsByQuizID(quizID)
}

func (u *Usecase) CreateAnswer(questionID int, answerText string, isCorrect bool) (int, error) {
	return u.p.CreateAnswer(questionID, answerText, isCorrect)
}

func (u *Usecase) GetAnswersByQuestionID(questionID int) ([]models.Answers, error) {
	return u.p.GetAnswersByQuestionID(questionID)
}

func (u *Usecase) SaveQuizRating(quizID int, rating models.Rating) error {
	return u.p.SaveRating(quizID, rating)
}

func (u *Usecase) GetAverageRating(quizID int) (float64, error) {
	return u.p.GetAverageRating(quizID)
}

func (u *Usecase) GetCommentsByQuizID(quizID int) ([]models.Rating, error) {
	return u.p.GetCommentsByQuizID(quizID)
}
