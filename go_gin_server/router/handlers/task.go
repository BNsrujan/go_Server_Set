package handlers

import (
	"fmt"
	"net/http"

	"github.com/BNsrujan/go_Server_Set.git/db"
	"github.com/gin-gonic/gin"
)

func SaveTask(c *gin.Context) {

	//marting
	var payload db.PostTaskReload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong the task was not added",
		})
		return
	}
	if payload.Title == "" {
		c.JSON(http.StatusOK, gin.H{"error": "Title error"})
		return
	}
	fmt.Println(payload)

	id, err := db.TaskRepository.SaveTaskQuery(payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"mess":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": true,
		"id":    id,
	})

	// var id int

	// query := `INSERT into tasks (title,description,status) VALUES ($1,$2,$3) RETURNING 	id;`

	// err := db.DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": true, "msg": err.Error(),
	// 	})
	// }
	// c.JSON(http.StatusOK, gin.H{"error": false, "Task_id": id})

	// body, err := c.GetRawData()

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "something went wrong",
	// 	})
	// }

	// fmt.Printf("data %v \n", body)

	// c.JSON(http.StatusOK, gin.H{
	// 	"task": "added",
	// })
}

// return tasks on success
// func ReadTask(ctx *gin.Context) {
// 	tasks, err := db.TaskRepository.ReadTask()

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"error": false, "msg": tasks,
// 	})
// }
