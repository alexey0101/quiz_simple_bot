package models

type Quiz struct {
	ID          int
	Title       string
	Description string
	Owner       int
	Question    *[]Question
}

type Question struct {
	ID           int
	Text         string
	QuizID       int
	QuestionText int
	Score        int
	AnswerOption *[]AnswerOption
}

type AnswerOption struct {
	ID         int
	Answer     string
	QuestionID int
	IsCorrect  bool
}

type QuizResponse struct {
	ID               int
	QuizID           int
	UserID           int
	Score            int
	QuestionResponse *[]QuestionResponse
}

type QuestionResponse struct {
	ID             int
	QuizResponseID int
	AnswerOption   int
}
