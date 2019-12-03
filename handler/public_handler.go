package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/letanthang/nc_student/db"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestDB(c echo.Context) error {
	return c.JSON(http.StatusOK, db.Test)
}

func GetAllStudents(c echo.Context) error {
	type Student struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		ClassName string `json:"class_name"`
		Email     string `json:"email"`
		Age       int    `json:"age"`
	}
	inputJson := `[{"first_name":"Tam","last_name":"Nguyen","age":100,"email":"tamnguyen@gmail.com"},{"first_name": "Hieu", "last_name": "Nguyen", "age":200,"email":"hieunguyen@gmail.com"}]`

	var students []Student
	json.Unmarshal([]byte(inputJson), &students)

	return c.JSON(http.StatusOK, students)
}

func AddStudent(c echo.Context) error {
	return c.JSON(http.StatusOK, db.Test)
}
