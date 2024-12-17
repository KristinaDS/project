package usecase

import (
	"backend/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// StartQuiz - старт викторины для пользователя
func (u *Usecase) StartQuiz(quizID int, userID int) (int, error) {
	// Начало викторины
	startedAt := time.Now().Format(time.RFC3339)

	// Создание записи о викторине
	quizParticipationID, err := u.p.StartQuiz(quizID, userID, startedAt)
	if err != nil {
		return 0, fmt.Errorf("failed to start quiz: %v", err)
	}

	return quizParticipationID, nil
}

// FinishQuiz - завершение викторины, подсчет баллов и сохранение результатов
func (u *Usecase) FinishQuizU(participationID int, answers []models.User_answers, completedAt string) (int, int, error) {
	// 1. Обновляем время завершения викторины
	err := u.p.FinishQuiz(participationID, completedAt)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to finish quiz: %v", err)
	}

	totalScore, elapsedTime, err := u.p.SaveAnswersAndGetScore(participationID, answers, completedAt)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to save answers and calculate score: %v", err)
	}

	return totalScore, elapsedTime, nil
}

func (u *Usecase) GetQuizParticipationByID(id int) (models.QuizParticipation, error) {
	// Получаем участие в викторине по ID
	quizParticipation, err := u.p.GetQuizParticipationByID(id)
	if err != nil {
		return models.QuizParticipation{}, fmt.Errorf("failed to get quiz participation: %v", err)
	}

	return quizParticipation, nil
}

// Метод для получения информации о викторине из quiz-service
func (u *Usecase) GetQuizInfo(quizID int) (models.Quiz, error) {
	// Формируем URL для запроса
	url := fmt.Sprintf("%s/quizzes/%d", u.quizServiceURL, quizID)

	// Отправляем GET запрос
	resp, err := http.Get(url)
	if err != nil {
		return models.Quiz{}, fmt.Errorf("failed to get quiz info: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return models.Quiz{}, fmt.Errorf("quiz service returned error: %v", resp.StatusCode)
	}

	// Парсим тело ответа
	var quiz models.Quiz
	if err := json.NewDecoder(resp.Body).Decode(&quiz); err != nil {
		return models.Quiz{}, fmt.Errorf("failed to parse quiz response: %v", err)
	}

	return quiz, nil
}

func (u *Usecase) SaveAnswersAndGetScore(participationID int, answers []models.User_answers, completedAt string) (int, int, error) {
	return u.p.SaveAnswersAndGetScore(participationID, answers, completedAt)
}

// GetAllParticipations - Получает все прохождения викторины
func (u *Usecase) GetAllParticipations(quizID int) ([]models.QuizParticipation, error) {
	// Используем Provider для получения всех прохождений
	participations, err := u.p.GetAllParticipations(quizID)
	if err != nil {
		return nil, err
	}
	return participations, nil
}
