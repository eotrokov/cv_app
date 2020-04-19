package queries

import (
	"cv/src/db"
	"cv/src/types"

	"github.com/graphql-go/graphql"
)

var getEmployeeById = &graphql.Field{
	Type:        types.EmployeeType,
	Description: "Get employee by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			return db.GetEmployeeById(id), nil
		}
		return nil, nil
	},
}

var getEmployees = &graphql.Field{
	Type:        graphql.NewList(types.EmployeeType),
	Description: "Get employees",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return db.GetEmployees(), nil
	},
}
