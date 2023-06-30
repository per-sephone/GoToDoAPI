// delete
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type listItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Status      string `json:"status"`
}

var list = []listItem{
	{ID: "1",
		Title:       "Finish project",
		Description: "Complete the API documentation",
		DueDate:     "2023-07-15",
		Status:      "In Progress"},
	{ID: "2",
		Title:       "Go grocery shopping",
		Description: "Buy fruits, vegetables, and milk",
		DueDate:     "2023-07-01",
		Status:      "Pending"},
	{ID: "3",
		Title:       "Read book",
		Description: "Finish reading 'The Great Gatsby'",
		DueDate:     "2023-07-10",
		Status:      "Completed"},
	{ID: "4",
		Title:       "Exercise",
		Description: "Go for a 30-minute jog",
		DueDate:     "2023-07-02",
		Status:      "In Progress"},
}

func main() {
	router := gin.Default()                    //initialize gin router
	router.GET("/list", getList)               //get request with endpoint name "list"
	router.GET("/list/:id", getItemByID)       //get request for specific item using ID
	router.POST("/list", postList)             //post request with endpoint name list
	router.PUT("/list/:id", editList)          //edit a list item by id number
	router.DELETE("/list/:id", deleteItemByID) //delete and item by id
	router.Run("localhost:8080")               //run the host
}

// get list items
func getList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, list)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range list {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func postList(c *gin.Context) {
	var newItem listItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	list = append(list, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func editList(c *gin.Context) {
	id := c.Param("id")
	var newItem listItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	for i, item := range list {
		if item.ID == id {
			//swap the item w new item
			list = replace(i, newItem)
			c.IndentedJSON(http.StatusCreated, newItem)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func replace(i int, newItem listItem) []listItem {
	front := list[:i]
	back := list[i+1:]
	front = append(front, newItem)
	front = append(front, back...)
	return front
}

func deleteItemByID(c *gin.Context) {
	id := c.Param("id")

	for i, item := range list {
		if item.ID == id {
			list = delete(i)
			c.IndentedJSON(http.StatusOK, list)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func delete(i int) []listItem {
	front := list[:i]
	back := list[i+1:]
	front = append(front, back...)
	return front
}
