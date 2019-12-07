package db

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type Error struct {
	Code int
	Msg  string
}

type StudentUpdateRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
