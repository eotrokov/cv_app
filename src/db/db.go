package db

import (
	"cv/src/models"
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*models.Employee)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func Connection() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",

		Addr: "localhost:5432",
	})
	return db
	//defer db.Close()
}

func CreateSchema() {
	db := Connection()
	defer db.Close()
	err := createSchema(db)
	if err != nil {
		panic(err)
	}
}

func ExampleDB_Model() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",

		Addr: "localhost:5432",
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	user1 := &models.Employee{
		LastName:  "admin",
		FirstName: "admin2@admin",
	}
	err = db.Insert(user1)
	if err != nil {
		panic(err)
	}

	err = db.Insert(&models.Employee{
		LastName:  "root",
		FirstName: "root2@root",
	})
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	employee := &models.Employee{Id: user1.Id}
	err = db.Select(employee)
	if err != nil {
		panic(err)
	}

	// Select all users.
	var employees []models.Employee
	err = db.Model(&employees).Select()
	if err != nil {
		panic(err)
	}

	// Select story and associated author in one query.
	// story := new(Story)
	// err = db.Model(story).
	// 	Relation("Author").
	// 	Where("story.id = ?", story1.Id).
	// 	Select()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(employee)
	fmt.Println(employees)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}
