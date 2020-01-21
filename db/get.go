package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllStudent() (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{} //map[string]interface{}
	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

func SearchStudentSimple(req StudentSearchRequest) (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// var filter bson.M
	// bs, err := json.Marshal(req)
	// err = json.Unmarshal(bs, &filter)
	// if err != nil {
	// 	log.Printf("marshal error: %v", err)
	// }

	filter := bson.M{}
	if req.ID != 0 {
		filter["id"] = req.ID
	}

	if req.LastName != "" {
		filter["last_name"] = req.LastName
	}

	if req.FirstName != "" {
		filter["first_name"] = req.FirstName
	}

	if req.ClassName != "" {
		filter["class_name"] = req.ClassName
	}
	if req.Age != 0 {
		filter["age"] = req.Age
	}

	fmt.Println(filter)

	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}
