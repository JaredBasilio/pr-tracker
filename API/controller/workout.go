package controller

import (
	"PR-Tracker/api/helper"
	"PR-Tracker/api/model"
	"net/http"
	"strconv"

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

	if workoutExists := ExistsWorkout(user.Workouts, newWorkout); workoutExists {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Workout already exists"})
		return
	}

	newWorkout.UserID = user.ID

	if savedWorkout, err := newWorkout.Save(); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error Saving Workout", "error": err.Error()})
	} else {
		context.IndentedJSON(http.StatusCreated, savedWorkout)
	}
}

func GetAllWorkouts(context *gin.Context) {
	if user, err := helper.CurrentUser(context); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": user.Workouts})
	}
}

func DeleteWorkout(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := context.Param("id")

	// Checks if we have a workout to add to
	if idParam == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no workout id found"})
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workout, err := model.FindWorkoutById(uint(id))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting current workout", "error": err.Error()})
		return
	}

	// Check if we own the workout so that we can add to it
	if workout.UserID != user.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user does not own workout"})
		return
	}

	if err := workout.Delete(); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error deleting workout", "error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "deleted workout"})
}

func ExistsWorkout(workouts []model.Workout, newWorkout model.Workout) bool {
	for _, w := range workouts {
		if w.Name == newWorkout.Name {
			return true
		}
	}
	return false
}
