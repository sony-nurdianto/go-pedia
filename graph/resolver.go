package graph

import "github.com/sony-nurdianto/go-pedia/graph/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver Exported
type Resolver struct {
	ProductRepo postgres.ProductRepo
	UserRepo    postgres.UserRepo
}
