package repository_test

import (
	"context"
	"database/sql"
	"indication/internal/domain"
	"indication/internal/repository"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_SetMeterIndicationByGUID(t *testing.T) {
	mockdb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockdb.Close()
	sqlxDB := sqlx.NewDb(mockdb, "postgres")
	repo := repository.NewNumLSRepository(sqlxDB)
	ctx := context.Background()
	guid := "123e4567-e89b-12d3-a456-426614174000"
	meter := float32(100.50)
	request := int64(12345)
	query := "UPDATE metering_devices"
	t.Run("Succeful Update", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(meter, guid).
			WillReturnResult(sqlmock.NewResult(1, 1))

		rowsAffected, err := repo.SetMeterIndicationByGUID(ctx, guid, meter, request)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), rowsAffected)
	})
	t.Run("No Row update", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(meter, guid).
			WillReturnResult(sqlmock.NewResult(1, 0))

		rowsAffected, err := repo.SetMeterIndicationByGUID(ctx, guid, meter, request)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), rowsAffected)
	})
}

func Test_GetObjectsByNumLS(t *testing.T) {
	mockdb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockdb.Close()

	sqlxDB := sqlx.NewDb(mockdb, "postgres")
	repo := repository.NewNumLSRepository(sqlxDB)
	ctx := context.Background()
	numLs := int64(11111)
	query := `SELECT ls.num_ls, 
	t.type, adr.address,
	md.day_night_type, 
	md.device_type, 
	md.last_metering,
	md.device_number,
	md.device_guid
	FROM ls_objects AS ls
	JOIN addresses as adr ON ls.address_id = adr.id
	JOIN object_types as t ON ls.object_type_id = t.id
	JOIN metering_devices md ON ls.id = md.ls_object_id
	WHERE ls.num_ls = $1
	`
	expectRow := domain.LS_Object{
		NumLS:        numLs,
		Type:         "Apartment",
		Address:      "123 Main Street",
		DayNightType: "Common",
		DeviceGuid:   "123e4567-e89b-12d3-a456-426614174000",
		DeviceType:   "CE 102 R5",
		LastMetering: 100.50,
		DeviceNumber: "XYZ1234567890",
	}
	rows := sqlmock.NewRows([]string{
		"num_ls", "type", "address",
		"day_night_type", "device_type",
		"last_metering", "device_number", "device_guid"}).
		AddRow(
			expectRow.NumLS, expectRow.Type, expectRow.Address,
			expectRow.DayNightType, expectRow.DeviceType,
			expectRow.LastMetering, expectRow.DeviceNumber, expectRow.DeviceGuid,
		)
	t.Run("Successful query", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(numLs).
			WillReturnRows(rows)
		row, err := repo.GetObjectsByNumLS(ctx, numLs)
		assert.NoError(t, err)
		assert.Len(t, row, 1)
		assert.Equal(t, expectRow, row[0])
	})
	t.Run("No Rows Found", func(t *testing.T) {
		emptyResult := sqlmock.NewRows([]string{
			"num_ls", "type", "address",
			"day_night_type", "device_type",
			"last_metering", "device_number", "device_guid",
		})
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(numLs).WillReturnRows(emptyResult)

		rows, err := repo.GetObjectsByNumLS(ctx, numLs)
		assert.NoError(t, err)
		assert.Empty(t, rows)
	})

	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(numLs).WillReturnError(sql.ErrConnDone)
		rows, err := repo.GetObjectsByNumLS(ctx, numLs)
		assert.Error(t, err)
		assert.Nil(t, rows)
	})
}
