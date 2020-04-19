package mutations

import (
	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new employee item
		http://localhost:8080/?query=mutation+_{create(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
		*/
		"create": create,
		/* Update employee by id
		   http://localhost:8080/?query=mutation+_{update(id:1,price:3.95){id,name,info,price}}
		*/
		"update": update,
		/* Delete employee by id
		   http://localhost:8080/?query=mutation+_{delete(id:1){id,name,info,price}}
		*/
		"delete": delete,
	},
})
