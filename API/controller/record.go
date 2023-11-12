package controller

import (
	"PR-Tracker/api/helper"
	"PR-Tracker/api/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddRecord(context *gin.Context) {
	var newRecord model.Record

	// Binds newRecord to JSON
	if err := context.BindJSON(&newRecord); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error binding json", "error": err.Error()})
		return
	}

	idParam := context.Param("id")

	// Checks if we have a workout to add to
	if idParam == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no workout id found"})
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gets Current User
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting current user", "error": err.Error()})
		return
	}

	// Gets Workout we want to add to
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

	newRecord.WorkoutID = workout.ID
	newRecord.Date = time.Now()

	savedRecord, err := newRecord.Save()

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error saving record", "error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusCreated, savedRecord)
}

func GetAllRecords(context *gin.Context) {
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

	// Check if we own the workout so that we can add to it
	if workout.UserID != user.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user does not own workout"})
		return
	}

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting current workout", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": workout.Records})
}
