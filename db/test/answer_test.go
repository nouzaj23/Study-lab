package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "study_lab/db/sqlc"
	"testing"
)

func createRandomAnswer(t *testing.T, question db.Question) db.Answer {
	arg := db.CreateAnswerParams{
		QuestionID: question.ID,
		Text:       RandomString(10),
		IsCorrect:  false,
	}

	answer, err := testQueries.CreateAnswer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, answer)

	require.Equal(t, arg.QuestionID, answer.QuestionID)
	require.Equal(t, arg.Text, answer.Text)
	require.Equal(t, arg.IsCorrect, answer.IsCorrect)

	require.NotZero(t, answer.ID)
	require.NotZero(t, answer.CreatedAt)

	return answer
}

func TestCreateAnswer(t *testing.T) {
	quiz := createRandomQuiz(t)
	question := createRandomQuestion(t, quiz)
	createRandomAnswer(t, question)
}

func TestGetAnswer(t *testing.T) {
	quiz := createRandomQuiz(t)
	question := createRandomQuestion(t, quiz)
	answer := createRandomAnswer(t, question)

	dbanswer, err := testQueries.GetAnswer(context.Background(), answer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dbanswer)

	require.Equal(t, answer.QuestionID, dbanswer.QuestionID)
	require.Equal(t, answer.Text, dbanswer.Text)
	require.Equal(t, answer.IsCorrect, dbanswer.IsCorrect)
}

func TestUpdateAnswer(t *testing.T) {
	quiz := createRandomQuiz(t)
	question := createRandomQuestion(t, quiz)
	answer := createRandomAnswer(t, question)

	arg := db.UpdateAnswerParams{
		ID:        answer.ID,
		Text:      RandomString(10),
		IsCorrect: answer.IsCorrect,
	}

	dbanswer, err := testQueries.UpdateAnswer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, dbanswer)

	require.Equal(t, answer.QuestionID, dbanswer.QuestionID)
	require.Equal(t, arg.Text, dbanswer.Text)
	require.Equal(t, answer.IsCorrect, dbanswer.IsCorrect)
}
