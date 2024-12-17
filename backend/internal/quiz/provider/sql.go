package provider

import (
	"backend/pkg/models"
	"database/sql"
	"fmt"
)

const (
	createQuizQuery      = `INSERT INTO quizzes (title, description) VALUES ($1, $2) RETURNING id`
	getAllQuizzesQuery   = `SELECT id, title, description FROM quizzes`
	getQuizByIDQuery     = `SELECT id, title, description FROM quizzes WHERE id = $1`
	createQuestionQuery  = `INSERT INTO questions (quiz_id, question_text) VALUES ($1, $2) RETURNING id`
	getQuestionsByQuizID = `SELECT id, question_text FROM questions WHERE quiz_id = $1`
	createAnswerQuery    = `INSERT INTO answers (question_id, answer_text, is_correct) VALUES ($1, $2, $3) RETURNING id`
	getAnswersByQuestion = `SELECT id, answer_text, is_correct FROM answers WHERE question_id = $1`
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

		questions, err := p.GetQuestionsByQuizID(quiz.ID)
		if err != nil {
			return nil, err
		}

		quiz.Questions = questions

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

	err := p.conn.QueryRow(getQuizByIDQuery, id).Scan(&quiz.ID, &quiz.Title, &quiz.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return quiz, nil
		}
		return quiz, err
	}

	questions, err := p.GetQuestionsByQuizID(id)
	if err != nil {
		return quiz, err
	}

	quiz.Questions = questions

	return quiz, nil
}

// CreateQuestion добавляет новый вопрос в викторину и возвращает его ID
func (p *Provider) CreateQuestion(quizID int, questionText string) (int, error) {
	var questionID int
	err := p.conn.QueryRow(createQuestionQuery, quizID, questionText).Scan(&questionID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert question: %v", err)
	}

	return questionID, nil
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
		if err := rows.Scan(&question.ID, &question.Question_text); err != nil {
			return nil, err
		}

		answer, err := p.GetAnswersByQuestionID(question.ID)
		if err != nil {
			return nil, err
		}

		if len(answer) > 0 {
			question.Answer = answer[0]
		}

		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (p *Provider) CreateAnswer(questionID int, answerText string, isCorrect bool) (int, error) {
	var answerID int
	err := p.conn.QueryRow(createAnswerQuery, questionID, answerText, isCorrect).Scan(&answerID)
	if err != nil {
		return 0, fmt.Errorf("failed to create answer: %v", err)
	}

	return answerID, nil
}

func (p *Provider) GetAnswersByQuestionID(questionID int) ([]models.Answers, error) {
	rows, err := p.conn.Query(getAnswersByQuestion, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var answers []models.Answers
	for rows.Next() {
		var answer models.Answers
		if err := rows.Scan(&answer.ID, &answer.Answer_text, &answer.Is_correct); err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return answers, nil
}

func (p *Provider) SaveRating(quizID int, rating models.Rating) error {
	query := `INSERT INTO ratings (quiz_id, score, comment) VALUES ($1, $2, $3)`
	_, err := p.conn.Exec(query, quizID, rating.Score, rating.Comment)
	if err != nil {
		return fmt.Errorf("could not insert rating: %v", err)
	}
	return nil
}

func (p *Provider) GetAverageRating(quizID int) (float64, error) {
	query := `SELECT AVG(score) AS average_rating FROM ratings WHERE quiz_id = $1`
	var averageRating float64
	err := p.conn.QueryRow(query, quizID).Scan(&averageRating)
	if err != nil {
		return 0, fmt.Errorf("could not calculate average rating: %v", err)
	}

	if averageRating == 0 {
		return 0, nil
	}

	return averageRating, nil
}

func (p *Provider) GetCommentsByQuizID(quizID int) ([]models.Rating, error) {
	query := `SELECT comment FROM ratings WHERE quiz_id = $1`
	rows, err := p.conn.Query(query, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Rating
	for rows.Next() {
		var comment models.Rating
		if err := rows.Scan(&comment.Comment); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
