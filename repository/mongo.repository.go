package repository

import (
	"context"
	"fmt"
	"go-graphql-mongodb-demo/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Save(user *model.User)
	FindAll() []*model.User
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "go-graphql-mongodb-demo"
	COLLECTION = "users"
)

func Init() UserRepository {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return &database{
		client: client,
	}
}

func (db *database) Save(user *model.User) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

}

func (db *database) FindAll() []*model.User {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var users []*model.User
	for cursor.Next(context.TODO()) {
		var user *model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
