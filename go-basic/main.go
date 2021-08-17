package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Binding from JSON
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main()  {
	fmt.Println("accc")
	r := gin.Default()
	r.POST("/hello", func(c *gin.Context) {
		var json User
		err := c.ShouldBindJSON(&json)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"errMsg":err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"username":json.Username,
			"password":json.Password,
			"age":json.Age,
		})
	})
	r.Run(":9090")
}

func add()  {
	
}
