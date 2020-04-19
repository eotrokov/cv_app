package mutations

import (
	"cv/src/models"
	"cv/src/store"
	"cv/src/types"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

var create = &graphql.Field{
	Type:        types.EmployeeType,
	Description: "Create new employee",
	Args: graphql.FieldConfigArgument{
		"firstName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"lastName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		rand.Seed(time.Now().UnixNano())
		employee := models.Employee{
			Id:        int(rand.Intn(100000)), // generate random ID
			FirstName: params.Args["firstName"].(string),
			LastName:  params.Args["lastName"].(string),
		}
		store.Employees = append(store.Employees, employee)
		return employee, nil
	},
}

var update = &graphql.Field{
	Type:        types.EmployeeType,
	Description: "Update employee by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"firstName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"lastName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)
		firstName, firstNameOk := params.Args["firstName"].(string)
		lastName, lastNameOk := params.Args["lastName"].(string)
		employee := models.Employee{}
		for i, p := range store.Employees {
			if int(id) == p.Id {
				if firstNameOk {
					store.Employees[i].FirstName = firstName
				}
				if lastNameOk {
					store.Employees[i].LastName = lastName
				}
				employee = store.Employees[i]
				break
			}
		}
		return employee, nil
	},
}

var delete = &graphql.Field{
	Type:        types.EmployeeType,
	Description: "Delete employee by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)
		employee := models.Employee{}
		for i, p := range store.Employees {
			if int(id) == p.Id {
				employee = store.Employees[i]
				// Remove from employee list
				store.Employees = append(store.Employees[:i], store.Employees[i+1:]...)
			}
		}
		return employee, nil
	},
}
