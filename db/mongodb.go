package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-student-service/dto"
	"os"
)

// ConnectDB connects to the mongodb database
func ConnectDB() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	return client, nil
}

type MongoDBClient struct {
	Client *mongo.Client
}

func (m *MongoDBClient) SaveStudent(student dto.Student) error {
	collection := m.Client.Database("yourDatabaseName").Collection("students")
	_, err := collection.InsertOne(context.TODO(), student)
	return err
}
