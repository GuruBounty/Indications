package repository_test

import (
	"context"
	"database/sql"
	"indication/internal/domain"
	"indication/internal/repository"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var (
	mockDBx *sqlx.DB
	mock    sqlmock.Sqlmock
	err     error
)

func TestMain(m *testing.M) {

	db, _, err := sqlmock.New()
	if err != nil {
		panic("Failed to create mock database: " + err.Error())
	}
	mockDBx = sqlx.NewDb(db, "sqlmock")

	// mockDB, mock, err = sqlmock.New()
	// if err != nil {
	// 	panic("Failde to create mock database: " + err.Error())
	// }
	defer mockDBx.Close()

	code := m.Run()
	os.Exit(code)

}

func TestUsers_GetByCredentials(t *testing.T) {
	mockdb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockdb.Close()
	sqlxDB := sqlx.NewDb(mockdb, "postgres")
	repo := repository.NewUsers(sqlxDB)
	ctx := context.Background()
	//users := repository.NewUsers(mockdb)
	email := "test@example.com"
	password := "testpassword123"
	inncorectemil := "wrongemail@example.com"
	inncorectpassword := "wrongpassword@example.com"

	expectedUser := domain.User{
		UserID: 123,
		Email:  email,
	}

	row := sqlmock.NewRows([]string{"id", "email"}).AddRow(expectedUser.UserID, expectedUser.Email)
	query := "SELECT id, email FROM users WHERE email= \\$1 AND password= \\$2"
	t.Run("Valid credentials", func(t *testing.T) {

		mock.ExpectQuery(query).WithArgs(email, password).WillReturnRows(row)
		user, err := repo.GetByCredentials(ctx, email, password)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfiled expectations: %s", err)
		}
	})
	t.Run("Ivalid credentials if email is wrong", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(inncorectemil, password).
			WillReturnError(sql.ErrNoRows)
		rows, err := repo.GetByCredentials(ctx, inncorectemil, password)
		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Empty(t, rows)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfiled expectations: %s", err)
		}
	})
	t.Run("Invaild credentials if password is incorrect", func(t *testing.T) {
		mock.ExpectQuery(query).WithArgs(email, inncorectpassword).
			WillReturnError(sql.ErrNoRows)
		rows, err := repo.GetByCredentials(ctx, email, inncorectpassword)
		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Empty(t, rows)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfiled expectations: %s", err)
		}
	})
}
