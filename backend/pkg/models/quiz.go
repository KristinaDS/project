package models

// Quiz представляет структуру данных для викторины
type Quiz struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Questions   []Questions `json:"questions"`
}

type Questions struct {
	ID            int    `json:"id"`
	Quiz_id       int    `json:"quiz_id"`
	Question_text string `json:"question_text"`
}

// Result представляет структуру данных для результата викторины
type Result struct {
	QuizID  int   `json:"quiz_id"`
	UserID  int   `json:"user_id"`
	Score   int   `json:"score"`
	Answers []int `json:"answers"`
}
