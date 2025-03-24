package repository

import (
	"context"
	"database/sql"
	"fmt"
	"indication/internal/domain"
	"time"
)

type MeterIndication struct {
	EventDateTime   time.Time
	IndicationValue float32
	RequestID       int
}

type MeterIndicationsHistoryLists struct {
	db *sql.DB
}

func NewMeterIndicationsHistoryLists(db *sql.DB) *MeterIndicationsHistoryLists {
	return &MeterIndicationsHistoryLists{db}
}

func (m *MeterIndicationsHistoryLists) CreateMeterIndicationsHistory(ctx context.Context, meterIndicationHistoryLists domain.MeterIndicationsHistoryLists) error {

	stm, err := m.db.Prepare("INSERT INTO meter_indications_history_lists(EventDateTime, IndicationValue, RequestId, UserId) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(meterIndicationHistoryLists.EventDateTime, meterIndicationHistoryLists.IndicationValue, meterIndicationHistoryLists.RequestID, meterIndicationHistoryLists.UserID)
	return err
}

func (m *MeterIndicationsHistoryLists) GetMeterIndicationsHistory(ctx context.Context, requestID int) (MeterIndication, error) {
	stm, err := m.db.Prepare("SELECT EventDateTime, IndicationValue, RequestID  FROM meter_indications_history_lists WHERE RequestId=$1")
	if err != nil {
		return MeterIndication{}, err
	}
	defer stm.Close()
	var indication MeterIndication
	err = stm.QueryRow(requestID).Scan(&indication.EventDateTime, &indication.IndicationValue, &indication.RequestID)
	if err != nil {
		if err != sql.ErrNoRows {
			return MeterIndication{}, fmt.Errorf("no indication found for request ID %d", requestID)

		}
		return MeterIndication{}, err
	}
	return indication, nil
}
