package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "study_lab/db/sqlc"
	"testing"
)

func createRandomQuestion(t *testing.T, quiz db.Quiz) db.Question {
	arg := db.CreateQuestionParams{
		QuizID: quiz.ID,
		Title:  RandomString(10) + "?",
	}

	question, err := testQueries.CreateQuestion(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, question)

	require.Equal(t, arg.QuizID, question.QuizID)
	require.Equal(t, arg.Title, question.Title)

	require.NotZero(t, question.ID)
	require.NotZero(t, question.CreatedAt)

	return question
}

func TestCreateQuestion(t *testing.T) {
	quiz := createRandomQuiz(t)
	createRandomQuestion(t, quiz)
}

func TestGetQuestion(t *testing.T) {
	quiz := createRandomQuiz(t)
	q := createRandomQuestion(t, quiz)
	dbq, err := testQueries.GetQuestion(context.Background(), q.ID)
	require.NoError(t, err)
	require.NotZero(t, dbq)

	require.Equal(t, q.ID, dbq.ID)
	require.Equal(t, q.Title, dbq.Title)
	require.Equal(t, q.QuizID, dbq.QuizID)
}

func TestUpdateQuestion(t *testing.T) {
	quiz := createRandomQuiz(t)
	q := createRandomQuestion(t, quiz)

	arg := db.UpdateQuestionParams{
		ID:    q.ID,
		Title: RandomString(10) + "?",
	}

	dbq, err := testQueries.UpdateQuestion(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, dbq)

	require.Equal(t, q.ID, dbq.ID)
	require.Equal(t, arg.Title, dbq.Title)
	require.Equal(t, q.QuizID, dbq.QuizID)
}
