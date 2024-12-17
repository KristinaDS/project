package models

type Quiz struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Questions   []Questions `json:"questions"`
}

type Questions struct {
	ID            int     `json:"id"`
	Quiz_id       int     `json:"quiz_id"`
	Question_text string  `json:"question_text"`
	Answer        Answers `json:"answer"`
}

type Answers struct {
	ID          int    `json:"id"`
	Answer_text string `json:"answer_text"`
	Is_correct  bool   `json:"is_correct"`
}

type Rating struct {
	Score   int    `json:"rating"`
	Comment string `json:"comment"`
}

type QuizParticipation struct {
	ID          int            `json:"id"`
	QuizID      int            `json:"quiz_id"`
	UserID      int            `json:"user_id"`
	StartedAt   string         `json:"started_at"`
	CompletedAt string         `json:"completed_at"`
	Answers     []User_answers `json:"answers"` // Список ответов пользователя
	TotalScore  int            `json:"total_score"`
	ElapsedTime int            `json:"elapsed_time"` // Время прохождения викторины в секундах
}

type User_answers struct { //ответы пользователя
	QuestionID int  `json:"question_id"`
	AnswerID   int  `json:"answer_id"`
	IsCorrect  bool `json:"is_correct"`
}

type AnalyticsReport struct {
	QuizReports []QuizReport `json:"quiz_reports"`
}

type QuizReport struct {
	QuizID         int     `json:"quiz_id"`
	AverageScore   float64 `json:"average_score"`
	AverageTime    float64 `json:"average_time"`
	CompletedCount int     `json:"completed_count"`
}

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	CreatedAt    string `json:"created_at"`
}
