package model

import (
	"PR-Tracker/api/database"

	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name    string `gorm:"type:text" json:"name"`
	UserID  uint
	Records []Record
}

func (workout *Workout) Save() (*Workout, error) {
	err := database.Database.Create(&workout).Error
	if err != nil {
		return &Workout{}, err
	}
	return workout, nil
}

func FindWorkoutById(id uint) (Workout, error) {
	var workout Workout
	err := database.Database.Preload("Records").Where("ID=?", id).Find(&workout).Error
	if err != nil {
		return Workout{}, err
	}
	return workout, nil
}
