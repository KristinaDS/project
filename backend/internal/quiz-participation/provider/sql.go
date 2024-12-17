package provider

import (
	"backend/pkg/models"
	"database/sql"
	"fmt"
)

// SQL-запросы
const (
	startQuizQuery            = `INSERT INTO quiz_participation (quiz_id, user_id, started_at) VALUES ($1, $2, $3) RETURNING id`
	saveAnswerQuery           = `INSERT INTO user_answers (participation_id, question_id, answer_id, is_correct) VALUES ($1, $2, $3, $4)`
	getTotalScoreQuery        = `SELECT SUM(CASE WHEN is_correct THEN 1 ELSE 0 END) FROM user_answers WHERE participation_id = $1`
	getElapsedTimeQuery       = `SELECT COALESCE(EXTRACT(EPOCH FROM (completed_at - started_at)), 0) FROM quiz_participation WHERE id = $1`
	finishQuizQuery           = `UPDATE quiz_participation SET completed_at = $2 WHERE id = $1`
	updateScoreAndTimeQuery   = `UPDATE quiz_participation SET total_score = $2, elapsed_time = $3 WHERE id = $1`
	getAllParticipationsQuery = `SELECT id, quiz_id, user_id, started_at, completed_at, total_score, elapsed_time FROM quiz_participation WHERE quiz_id = $1`
)

// StartQuiz сохраняет информацию о начале викторины и возвращает ID записи
func (p *Provider) StartQuiz(quizID int, userID int, startedAt string) (int, error) {
	var participationID int
	err := p.conn.QueryRow(startQuizQuery, quizID, userID, startedAt).Scan(&participationID)
	if err != nil {
		return 0, fmt.Errorf("failed to start quiz: %v", err)
	}
	return participationID, nil
}

// SaveAnswersAndGetScore - сохраняет ответы и вычисляет баллы
func (p *Provider) SaveAnswersAndGetScore(participationID int, answers []models.User_answers, completedAt string) (int, int, error) {
	// 1. Сначала обновляем время завершения викторины
	err := p.FinishQuiz(participationID, completedAt)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to finish quiz: %v", err)
	}

	// 2. Сохраняем ответы пользователя
	for _, answer := range answers {
		_, err := p.conn.Exec(saveAnswerQuery, participationID, answer.QuestionID, answer.AnswerID, answer.IsCorrect)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to save answer: %v", err)
		}
	}

	// 3. Получаем общий балл викторины
	var totalScore int
	err = p.conn.QueryRow(getTotalScoreQuery, participationID).Scan(&totalScore)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to calculate total score: %v", err)
	}

	// 4. Теперь вычисляем время прохождения
	var elapsedTime float64
	err = p.conn.QueryRow(getElapsedTimeQuery, participationID).Scan(&elapsedTime)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to calculate elapsed time: %v", err)
	}

	// 5. Обновляем поля total_score и elapsed_time
	_, err = p.conn.Exec(updateScoreAndTimeQuery, participationID, totalScore, int(elapsedTime))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to update total_score and elapsed_time: %v", err)
	}

	return totalScore, int(elapsedTime), nil
}

// FinishQuiz - завершение викторины
func (p *Provider) FinishQuiz(participationID int, completedAt string) error {
	_, err := p.conn.Exec(finishQuizQuery, participationID, completedAt)
	if err != nil {
		return fmt.Errorf("failed to finish quiz: %v", err)
	}
	return nil
}

func (p *Provider) GetQuizParticipationByID(id int) (models.QuizParticipation, error) {
	var participation models.QuizParticipation

	// Выполняем SQL-запрос для получения данных о прохождении викторины
	query := `
		SELECT id, quiz_id, user_id, started_at, completed_at, total_score, elapsed_time
		FROM quiz_participation
		WHERE id = $1
	`
	err := p.conn.QueryRow(query, id).Scan(&participation.ID, &participation.QuizID, &participation.UserID, &participation.StartedAt, &participation.CompletedAt, &participation.TotalScore, &participation.ElapsedTime)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если записи не существует
			return participation, nil
		}
		return participation, fmt.Errorf("failed to query participation: %v", err)
	}

	// Получаем список ответов для этого участия
	answersQuery := `
		SELECT question_id, answer_id, is_correct
		FROM answers_user
		WHERE participation_id = $1
	`
	rows, err := p.conn.Query(answersQuery, participation.ID)
	if err != nil {
		return participation, fmt.Errorf("failed to query answers: %v", err)
	}
	defer rows.Close()

	// Считываем ответы
	for rows.Next() {
		var answer models.User_answers
		if err := rows.Scan(&answer.QuestionID, &answer.AnswerID, &answer.IsCorrect); err != nil {
			return participation, fmt.Errorf("failed to scan answer: %v", err)
		}
		participation.Answers = append(participation.Answers, answer)
	}

	// Проверяем на наличие ошибок при обходе строк
	if err := rows.Err(); err != nil {
		return participation, fmt.Errorf("error iterating over answers: %v", err)
	}

	return participation, nil
}

func (p *Provider) GetQuizInfo(quizID int) (models.Quiz, error) {
	var quiz models.Quiz
	query := "SELECT id, title, description FROM quizzes WHERE id = $1"
	err := p.conn.Get(&quiz, query, quizID)

	if err != nil {
		return models.Quiz{}, fmt.Errorf("failed to get quiz info: %v", err)
	}

	return quiz, nil
}

func (p *Provider) GetAllParticipations(quizID int) ([]models.QuizParticipation, error) {
	var participations []models.QuizParticipation

	rows, err := p.conn.Query(getAllParticipationsQuery, quizID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch participations: %v", err)
	}
	defer rows.Close()

	// Чтение всех строк из результата
	for rows.Next() {
		var participation models.QuizParticipation
		if err := rows.Scan(&participation.ID, &participation.QuizID, &participation.UserID, &participation.StartedAt, &participation.CompletedAt, &participation.TotalScore, &participation.ElapsedTime); err != nil {
			return nil, fmt.Errorf("failed to scan participation: %v", err)
		}
		participations = append(participations, participation)
	}

	// Проверяем на наличие ошибок после завершения перебора строк
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate rows: %v", err)
	}

	return participations, nil
}

func (p *Provider) GetElapsedTime(participationID int) (int, error) {
	var elapsedTime float64
	query := `SELECT EXTRACT(EPOCH FROM (completed_at - started_at)) FROM quiz_participation WHERE id = $1`
	err := p.conn.QueryRow(query, participationID).Scan(&elapsedTime)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate elapsed time: %v", err)
	}
	return int(elapsedTime), nil
}
