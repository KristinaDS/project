package provider

import (
	"backend/pkg/models"
	"database/sql"
)

// SQL-запросы и логика работы с базой данных для викторин
const (
	createQuizQuery      = `INSERT INTO quizzes (title, description) VALUES ($1, $2) RETURNING id`
	getAllQuizzesQuery   = `SELECT id, title, description FROM quizzes`
	getQuizByIDQuery     = `SELECT id, title, description FROM quizzes WHERE id = $1`
	updateQuizQuery      = `UPDATE quizzes SET title = $1, description = $2 WHERE id = $3`
	deleteQuizQuery      = `DELETE FROM quizzes WHERE id = $1`
	createQuestionQuery  = `INSERT INTO questions (quiz_id, question_text) VALUES ($1, $2)`
	getQuestionsByQuizID = `SELECT id, question_text FROM questions WHERE quiz_id = $1`
)

// CreateQuiz создает новую викторину в базе данных
func (p *Provider) CreateQuiz(quiz models.Quiz) (int, error) {
	var quizID int
	err := p.conn.QueryRow(createQuizQuery, quiz.Title, quiz.Description).Scan(&quizID)
	if err != nil {
		return 0, err
	}
	return quizID, nil
}

// GetAllQuizzes возвращает все викторины из базы данных
func (p *Provider) GetAllQuizzes() ([]models.Quiz, error) {
	rows, err := p.conn.Query(getAllQuizzesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []models.Quiz
	for rows.Next() {
		var quiz models.Quiz
		if err := rows.Scan(&quiz.ID, &quiz.Title, &quiz.Description); err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quizzes, nil
}

// GetQuizByID возвращает викторину по ее ID
func (p *Provider) GetQuizByID(id int) (models.Quiz, error) {
	var quiz models.Quiz

	// Получаем викторину по ID
	err := p.conn.QueryRow(getQuizByIDQuery, id).Scan(&quiz.ID, &quiz.Title, &quiz.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return quiz, nil // Викторина не найдена, возвращаем пустой объект
		}
		return quiz, err
	}

	// Теперь получаем все вопросы для этой викторины
	questions, err := p.GetQuestionsByQuizID(id)
	if err != nil {
		return quiz, err
	}

	quiz.Questions = questions // Добавляем вопросы в объект викторины

	return quiz, nil
}

// UpdateQuiz обновляет информацию о викторине
func (p *Provider) UpdateQuiz(id int, quiz models.Quiz) error {
	_, err := p.conn.Exec(updateQuizQuery, quiz.Title, quiz.Description, id)
	return err
}

// DeleteQuiz удаляет викторину по ее ID
func (p *Provider) DeleteQuiz(id int) error {
	_, err := p.conn.Exec(deleteQuizQuery, id)
	return err
}

// CreateQuestion добавляет новый вопрос в викторину
func (p *Provider) CreateQuestion(quizID int, questionText string) error {
	_, err := p.conn.Exec(createQuestionQuery, quizID, questionText)
	return err
}

// GetQuestionsByQuizID возвращает все вопросы для викторины по ID
func (p *Provider) GetQuestionsByQuizID(quizID int) ([]models.Questions, error) {
	rows, err := p.conn.Query(getQuestionsByQuizID, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Questions
	for rows.Next() {
		var question models.Questions
		if err := rows.Scan(&question); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}
