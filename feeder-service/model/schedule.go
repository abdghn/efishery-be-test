package model

import "time"

type Schedule struct {
	ID            int       `json:"id" gorm:"primary_key"`
	Uuid          string    `json:"uuid" gorm:"unique"`
	PondUuid      string    `json:"pond_uuid" gorm:"column:pond_uuid"`
	TimeStart     time.Time `json:"time_start"`
	TimeEnd       time.Time `json:"time_end"`
	DurationRun   int       `json:"duration_run"`
	DurationPause int       `json:"duration_pause"`
	ScheduleType  string    `json:"schedule_type"`
	CreatedAt     time.Time `json:"created_at"`
}
