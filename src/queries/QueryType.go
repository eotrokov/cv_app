package queries

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single employee by id
			   http://localhost:8080/?query={employee(id:1){name}}
			*/
			"employee": getEmployeeById,
			/* Get (read) employees
			   http://localhost:8080/?query={employees{id,name}}
			*/
			"employees": getEmployees,
		},
	},
)
