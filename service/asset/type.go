package asset

import "time"

type (
	Asset struct {
		ID              string    `gorm:"column:id;primary_key;default:generate_uuid_v4()"`
		Name            string    `gorm:"column:name"`
		Type            string    `gorm:"column:type"`
		Value           float64   `gorm:"column:value"`
		AcquisitionDate time.Time `gorm:"column:acquisition_date;type:date"`
	}
)
