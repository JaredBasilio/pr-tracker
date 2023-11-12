package controller

import (
	"PR-Tracker/api/helper"
	"PR-Tracker/api/model"
	"fmt"
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

	id, err := strconv.ParseUint(idParam, 10, 64)

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
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting workout", "error": err.Error()})
		return
	}

	// Check if we own the workout so that we can add to it
	if workout.UserID != user.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user does not have rights to workout"})
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

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting workout", "error": err.Error()})
		return
	}

	// Check if we own the workout so that we can add to it
	if workout.UserID != user.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user does not have rights to workout"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": workout.Records})
}

func DeleteRecord(context *gin.Context) {
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

	// Gets Current User
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting current user", "error": err.Error()})
		return
	}

	// Gets Workout we want to remove from
	workout, err := model.FindWorkoutById(uint(id))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting workout", "error": err.Error()})
		return
	}

	// Check if we own the workout so that we can remove from it
	if workout.UserID != user.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user does not have rights to workout, workout not owned by user"})
		return
	}

	recordIdParam := context.Param("recordId")

	if recordIdParam == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no record id found"})
		return
	}

	recordId, err := strconv.ParseUint(recordIdParam, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := model.FindRecordById(uint(recordId))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting record", "error": err.Error()})
		return
	}

	if record.WorkoutID != workout.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user does not have rights to workout %d, record %d not apart of workout", workout.ID, record.WorkoutID)})
		return
	}

	err = record.Delete()

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error deleting record", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "removed record from workout"})
}
