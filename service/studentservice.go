package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-student-service/dto"
	"golang.org/x/net/context"
)

type MongoDBClient struct {
	Client *mongo.Client
}

type StudentService struct {
	MongoDBClient *MongoDBClient
}

func (m *MongoDBClient) SaveStudent(student dto.Student) error {
	collection := m.Client.Database("yourDatabaseName").Collection("students")
	_, err := collection.InsertOne(context.TODO(), student)
	return err
}
