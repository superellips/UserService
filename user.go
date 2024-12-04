package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name"`
}

func PostUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	var db MongoDb
	_, err := db.readByName(&newUser.Name)
	if err == nil {
		c.IndentedJSON(http.StatusNotAcceptable, nil)
		return
	}
	db.create(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
func GetUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		return
	}
	var db MongoDb
	user, err := db.read(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	var db MongoDb
	user, err := db.readByName(&name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func PutUser(c *gin.Context) {
	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var db MongoDb
	result, err := db.update(&updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, result)
}
func DeleteUser(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not a valid id"})
		return
	}
	var db MongoDb
	deletedUser, err := db.delete(&userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not able to remove user"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, deletedUser)
}
