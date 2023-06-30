// post
// get
// put
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
	router := gin.Default()      //initialize gin router
	router.GET("/list", getList) //get request with endpoint name "list"
	router.Run("localhost:8080") //run the host
}

// get list items
func getList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, list)
}
