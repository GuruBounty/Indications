package domain

import "time"

type MeterIndicationsHistoryLists struct {
	MeterIndicationsHistoryListsID int       `json:"id"`
	EventDateTime                  time.Time `json:"event_date_time"`
	IndicationValue                float32   `json:"indication_value"`
	RequestID                      int       `json:"request_id"`
	UserID                         int       `json:"user_id"`
}
