package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Data 1": "Hi",
	})
}

func postGin(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"Data 2": "Finish Post",
	})
}

// Trả về chính xác số id
func getID(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id:": id,
	})
}

func getName(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "Guest")
	ctx.JSON(http.StatusOK, gin.H{
		"Name:":   name,
		"Message": "Hello " + name,
	})

}

func main() {
	r := gin.Default()

	r.GET("/Get", getGin)
	r.POST("/Post", postGin)
	r.GET("/GetID/:id", getID) // Get ID after folder add /YOURID
	r.GET("/Getname", getName) // Get name after folder add ?name=XXX

	r.Run(":333")

}
