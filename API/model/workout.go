package model

import (
	"PR-Tracker/api/database"

	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	Name        string `gorm:"type:text" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint
	Records     []Record
}

func (workout *Workout) Save() (*Workout, error) {
	err := database.Database.Create(&workout).Error
	if err != nil {
		return &Workout{}, err
	}
	return workout, nil
}

func (workout *Workout) Delete() error {
	err := database.Database.Delete(workout).Error
	return err
}

func (workout *Workout) Update(id uint, update Workout) error {
	err := database.Database.Model(&Workout{}).Where("ID=?", id).Updates(update).Error
	return err
}

func FindWorkoutById(id uint) (Workout, error) {
	var workout Workout
	err := database.Database.Preload("Records").Where("ID=?", id).Find(&workout).Error
	if err != nil {
		return Workout{}, err
	}
	return workout, nil
}
