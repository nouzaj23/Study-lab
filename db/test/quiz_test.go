package db

import (
	"context"
	"database/sql"
	db "study_lab/db/sqlc"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomQuiz(t *testing.T) db.Quiz {
	name := RandomString(8)

	quiz, err := testQueries.CreateQuiz(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, quiz)

	require.Equal(t, name, quiz.Name)
	require.NotZero(t, quiz.ID)
	require.NotZero(t, quiz.CreatedAt)

	return quiz
}

func TestCreateQuiz(t *testing.T) {
	createRandomQuiz(t)
}

func TestGetQuiz(t *testing.T) {
	quiz := createRandomQuiz(t)
	dbquiz, err := testQueries.GetQuiz(context.Background(), quiz.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dbquiz)

	require.Equal(t, quiz.ID, dbquiz.ID)
	require.Equal(t, quiz.Name, dbquiz.Name)
}

func TestUpdateQuiz(t *testing.T) {
	quiz := createRandomQuiz(t)

	arg := db.UpdateQuizParams{
		ID:   quiz.ID,
		Name: RandomString(8),
	}

	dbquiz, err := testQueries.UpdateQuiz(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, dbquiz)

	require.Equal(t, quiz.ID, dbquiz.ID)
	require.Equal(t, arg.Name, dbquiz.Name)
}

func TestDeleteQuiz(t *testing.T) {
	quiz := createRandomQuiz(t)
	err := testQueries.DeleteQuiz(context.Background(), quiz.ID)
	require.NoError(t, err)

	dbquiz, err := testQueries.GetQuiz(context.Background(), quiz.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, dbquiz)
}

func TestListQuizzes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomQuiz(t)
	}

	arg := db.ListQuizzesParams{
		Limit:  5,
		Offset: 5,
	}

	quizzes, err := testQueries.ListQuizzes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, quizzes, 5)

	for _, quiz := range quizzes {
		require.NotEmpty(t, quiz)
	}
}
