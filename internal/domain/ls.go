package domain

import "errors"

var (
	ErrLSNotFound = errors.New("ls not found")
)

type LS struct {
	NumLS        int64          `json:"num_ls"`
	Address      string         `json:"address"`
	Type         string         `json:"type"`
	MeterObjects []MeterObjects `json:"mesuaringObject"`
}

type MeterObjects struct {
	DayNightType string  `json:"dayNightType"`
	DeviceNumber string  `json:"deviceNumber"`
	DeviceType   string  `json:"deviceType"`
	DeviceGuid   string  `json:"deviceGuid"`
	LastMetering float32 `json:"lastMetering"`
}
