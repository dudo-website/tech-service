package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dudo/go_hello_world/graph/generated"
	"dudo/go_hello_world/graph/model"
	"fmt"
	"math/rand"
)

func (r *mutationResolver) CreateTech(ctx context.Context, input model.NewTech) (*model.Technology, error) {
	tech := &model.Technology{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Intn(100)),
	}
	r.technologies = append(r.technologies, tech)
	return tech, nil
}

func (r *queryResolver) Technologies(ctx context.Context) ([]*model.Technology, error) {
	return r.technologies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
