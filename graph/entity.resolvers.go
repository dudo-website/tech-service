package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dudo/go_hello_world/graph/generated"
	"dudo/go_hello_world/graph/model"
	"fmt"
)

func (r *entityResolver) FindTechnologyByID(ctx context.Context, id string) (*model.Technology, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
