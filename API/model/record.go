package model

import (
	"PR-Tracker/api/database"
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	WorkoutID     uint
	Date          time.Time `gorm:"type:time" json:"date"`
	Weight        int       `gorm:"type:int" json:"weight"`
	Notes         string    `gorm:"type:text" json:"notes"`
	UnitOfMeasure string    `gorm:"type:text" json:"unit_of_measurement"`
}

func (record *Record) Save() (*Record, error) {
	err := database.Database.Create(&record).Error
	if err != nil {
		return &Record{}, err
	}
	return record, nil
}

func (record *Record) Delete() error {
	err := database.Database.Delete(record).Error
	return err
}

func FindRecordById(id uint) (Record, error) {
	var record Record
	err := database.Database.Where("ID=?", id).Find(&record).Error
	if err != nil {
		return Record{}, err
	}
	return record, nil
}
