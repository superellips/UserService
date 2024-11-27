package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	Uri string
}

func (p *MongoDb) create(u *User) *User {
	coll := p.userCollection()
	result, err := coll.InsertOne(
		context.TODO(),
		u)
	if err != nil {
		panic(err)
	}
	u.Id = result.InsertedID.(primitive.ObjectID)
	return u
}

func (p *MongoDb) read(i *primitive.ObjectID) *User {
	coll := p.userCollection()
	filter := bson.D{{Key: "_id", Value: i}}
	var result User
	err := coll.FindOne(
		context.TODO(),
		filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (p *MongoDb) readByName(n *string) (*User, error) {
	coll := p.userCollection()
	filter := bson.D{{Key: "name", Value: n}}
	var result User
	err := coll.FindOne(
		context.TODO(),
		filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *MongoDb) update(u *User) *User {
	coll := p.userCollection()
	filter := bson.D{{Key: "_id", Value: u.Id}}
	var result User
	err := coll.FindOneAndReplace(
		context.TODO(),
		filter,
		u).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (p *MongoDb) delete(i *primitive.ObjectID) *User {
	coll := p.userCollection()
	filter := bson.D{{Key: "_id", Value: i}}
	var deletedUser User
	err := coll.FindOneAndDelete(
		context.TODO(),
		filter).Decode(&deletedUser)
	if err != nil {
		panic(err)
	}
	return &deletedUser
}

func (p *MongoDb) userCollection() *mongo.Collection {
	p.Uri = os.Getenv("GB_CONSTRING")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(p.Uri))
	if err != nil {
		panic(err)
	}
	return client.Database("user-service").Collection("users")
}
