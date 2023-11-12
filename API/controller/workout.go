package controller

import (
	"PR-Tracker/api/helper"
	"PR-Tracker/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWorkout(context *gin.Context) {
	var newWorkout model.Workout
	if err := context.BindJSON(&newWorkout); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error Creating Workout", "error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error Creating Workout under User", "error": err.Error()})
		return
	}

	for _, w := range user.Workouts {
		if w.Name == newWorkout.Name {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Workout already exists"})
			return
		}
	}

	newWorkout.UserID = user.ID

	savedWorkout, err := newWorkout.Save()

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error Saving Workout", "error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusCreated, savedWorkout)
}

func GetAllWorkouts(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Workouts})
}
