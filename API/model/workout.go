package model

import (
	"PR-Tracker/api/database"
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name            string `gorm:"type:text" json:"name"`
	UserID          uint
	PersonalRecords []PersonalRecord
}

type PersonalRecord struct {
	gorm.Model
	WorkoutID     uint
	Date          time.Time `gorm:"type:time" json:"date"`
	Weight        int       `gorm:"type:int" json:"weight"`
	Notes         string    `gorm:"type:text" json:"notes"`
	UnitOfMeasure string    `gorm:"type:text" json:"unit_of_measurement"`
}

func (workout *Workout) Save() (*Workout, error) {
	err := database.Database.Create(&workout).Error
	if err != nil {
		return &Workout{}, err
	}
	return workout, nil
}
