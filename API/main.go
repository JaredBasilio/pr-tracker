package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/google/uuid"
	"errors"
	// "fmt"
)

type workout struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type pr struct {
	ID string `json:"id"`
	WorkoutID string `json:"workoutID`
	Date time.Time `json:"date"`
	Record int `json:"record"`
	Notes string `json:"notes"`
}

var workouts []workout

func createWorkout(c *gin.Context) {
	var newWorkout workout

	if err := c.BindJSON(&newWorkout); err != nil {
		return
	}

	// Check for duplicate workout name
	for _, w := range workouts {
		if w.Name == newWorkout.Name {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Workout already exists"})
			return
		}
	}

	newWorkout.ID = uuid.New().String()

	workouts = append(workouts, newWorkout)
	c.IndentedJSON(http.StatusCreated, newWorkout)
}

func getWorkouts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workouts)
}

// doesnt work
func updateWorkout(c *gin.Context) {
	id := c.Param("workoutId")
	workoutToChange, err := getWorkoutById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Workout not found."})
		return
	}

	newName, ok := c.GetQuery("newName")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing newName query parameter."})
		return
	}

	workoutToChange.Name = newName
	c.IndentedJSON(http.StatusOK, workoutToChange)
}

func workoutById(c *gin.Context) {
	id := c.Param("workoutId")
	workout, err := getWorkoutById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Workout not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, workout)
}

func getWorkoutById(id string) (*workout, error) {
	for i, w := range workouts {
		if w.ID == id {
			return &workouts[i], nil
		}
	}

	return nil, errors.New("workout not found")
}

func deleteWorkout(c *gin.Context) {
	id := c.Param("workoutId")

	for i, w := range workouts {
		if w.ID == id {
			workouts = append(workouts[:i], workouts[i + 1:]...)
			break
		}
	}

	// TODO: delete all occurrances of a PR where id == parentID

	c.Status(http.StatusNoContent)
}

func addPR(c *gin.context) {
	id := c.Param("workoutId")
}

func main() {
	router := gin.Default()

	router.POST("/workouts", createWorkout)
	router.GET("/workouts", getWorkouts)
	router.GET("/workouts/:workoutId", workoutById)
	router.PUT("/workouts/:workoutId/updateWorkout", updateWorkout)
	router.DELETE("/workouts/:workoutId/deleteWorkout", deleteWorkout)

	// router.POST("/workouts/:workoutId/addPR", addPR)
	// r.PUT("/:workoutId/:prId", updatePR)
	// r.DELETE("/:workoutId/:prId", deletePR)

	router.Run("localhost:8080")
}