package db

import (
	"cv/src/models"
)

func GetEmployees() []models.Employee {
	db := Connection()
	defer db.Close()
	var employees []models.Employee
	err := db.Model(&employees).Select()
	if err != nil {
		panic(err)
	}
	return employees
}

func GetEmployeeById(id int) models.Employee {
	db := Connection()
	defer db.Close()
	employee := &models.Employee{Id: id}
	err := db.Select(employee)
	if err != nil {
		panic(err)
	}
	return *employee
}
