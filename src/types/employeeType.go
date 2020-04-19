package types

import "github.com/graphql-go/graphql"

var EmployeeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
