package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-example/graph/generated"
	"graphql-example/graph/model"
)

// Programmers is the resolver for the programmers field.
func (r *queryResolver) Programmers(ctx context.Context, skill string) ([]*model.Programmer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
