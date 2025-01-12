package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "github.com/manoelsslima/fc-13-GraphQL/internal/database"

type Resolver struct{
	// injects database dependency
	CategoryDB *database.Category // now, adjust schema.resolves.go to access this
}
