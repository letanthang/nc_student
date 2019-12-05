package db

import (
	"context"
	"time"
)

func AddStudent(student *Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, student)
	return res, err
}
