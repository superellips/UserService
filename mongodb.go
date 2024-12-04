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
	Uri  string
	Coll *mongo.Collection
}

func (p *MongoDb) create(u *User) (*User, error) {
	if err := p.getUserCollection(); err != nil {
		return nil, err
	}
	result, err := p.Coll.InsertOne(
		context.TODO(),
		u)
	if err != nil {
		return nil, err
	}
	u.Id = result.InsertedID.(primitive.ObjectID)
	return u, nil
}

func (p *MongoDb) read(i *primitive.ObjectID) (*User, error) {
	if err := p.getUserCollection(); err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: i}}
	var result User
	err := p.Coll.FindOne(
		context.TODO(),
		filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *MongoDb) readByName(n *string) (*User, error) {
	if err := p.getUserCollection(); err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "name", Value: n}}
	var result User
	err := p.Coll.FindOne(
		context.TODO(),
		filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *MongoDb) update(u *User) (*User, error) {
	if err := p.getUserCollection(); err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: u.Id}}
	var result User
	err := p.Coll.FindOneAndReplace(
		context.TODO(),
		filter,
		u).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *MongoDb) delete(i *primitive.ObjectID) (*User, error) {
	if err := p.getUserCollection(); err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: i}}
	var deletedUser User
	err := p.Coll.FindOneAndDelete(context.TODO(), filter).Decode(&deletedUser)
	if err != nil {
		return nil, err
	}
	return &deletedUser, nil
}

func (p *MongoDb) getUserCollection() error {
	p.Uri = os.Getenv("GB_CONSTRING")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(p.Uri))
	if err != nil {
		return err
	}
	p.Coll = client.Database("user-service").Collection("users")
	return nil
}
