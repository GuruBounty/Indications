package domain

import "errors"

var ErrLSObjectNotFound = errors.New("no rows in result set")

type LS_Object struct {
	NumLS        int64   `db:"num_ls" json:"num_ls"`
	Address      string  `db:"address" json:"address"`
	Type         string  `db:"type" json:"type"`
	DayNightType string  `db:"day_night_type" json:"day_night_type"`
	DeviceGuid   string  `db:"device_guid" json:"device_guid"`
	DeviceType   string  `db:"device_type" json:"device_type"`
	LastMetering float32 `db:"last_metering" json:"last_metering"`
	DeviceNumber string  `db:"device_number" json:"device_number"`
}
