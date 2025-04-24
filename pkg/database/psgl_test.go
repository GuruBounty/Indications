package database_test

import (
	"database/sql"
	"fmt"
	"indication/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestNewPostgresConnection_ValidInfo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database, error: %v", err)
	}
	defer mockDB.Close()

	orginalOpen := database.SqlOpen
	defer func() { database.SqlOpen = orginalOpen }()

	database.SqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		return mockDB, nil
	}
	mock.ExpectPing().WillReturnError(nil)

	info := database.ConnectionInfo{
		Host:     "localhost",
		Port:     5555,
		Username: "postgres",
		Password: "321",
		DBName:   "Indications",
		SSLMode:  "disable",
	}
	db, err := database.NewPostgresConnection(info)

	assert.NoError(t, err)
	assert.NotNil(t, db)
	assert.Equal(t, mockDB, db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfufilled expectations: %s", err)
	}

}

func TestNewPostgresConnection_Failure(t *testing.T) {
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Failed to create mock database, error: %v", err)
	}
	defer mockDB.Close()

	orginalOpen := database.SqlOpen
	defer func() { database.SqlOpen = orginalOpen }()

	database.SqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		return mockDB, nil
	}
	mock.ExpectPing().WillReturnError(fmt.Errorf("ping error"))

	info := database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "testuser",
		Password: "testpassword",
		DBName:   "testdb",
		SSLMode:  "disable",
	}
	db, err := database.NewPostgresConnection(info)
	assert.Error(t, err)
	assert.Nil(t, db)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfufilled expectations: %s", err)
	}
}
