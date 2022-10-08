package model

type PondFeeder struct {
	ID         int    `json:"id" gorm:"primary_key"`
	PondUuid   string `json:"pond_uuid" gorm:"column:pond_uuid"`
	FeederUuid string `json:"feeder_uuid" gorm:"column:feeder_uuid"`
}
