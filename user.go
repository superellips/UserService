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
	user := db.read(&id)
	c.IndentedJSON(http.StatusOK, user)
}
func PutUser(c *gin.Context) {
	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}
	var db MongoDb
	updatedUser = *db.update(&updatedUser)
	c.IndentedJSON(http.StatusAccepted, updatedUser)
}
func DeleteUser(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		return
	}
	var db MongoDb
	deletedUser := *db.delete(&userId)
	c.IndentedJSON(http.StatusAccepted, deletedUser)
}
