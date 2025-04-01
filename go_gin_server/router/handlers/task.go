package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostTaskReload struct {
	Title     string `json:"title" binding:required`
	Description string `json:"description" binding:required`
	Status      string `json:"Status binding:required`
}

func SaveTask(c *gin.Context) {

	//marting
	var payload PostTaskReload

	err := c.ShouldBindBodyWithJSON(&payload)
	if err != nil {
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
	// var id int

	// quiry := `INSERT into tasks (title,description,status) VALUE ($1,$2,$3)`

	c.JSON(http.StatusOK, gin.H{"error": false, "payload": payload})

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
