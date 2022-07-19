package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"graphql-example/db"
)

type Resolver struct {
	DB db.DB
}
