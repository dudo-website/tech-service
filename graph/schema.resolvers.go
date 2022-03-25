package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dudo/tech_service/graph/generated"
	"dudo/tech_service/graph/model"
	"fmt"
	"math/rand"
)

func (r *queryResolver) Technologies(ctx context.Context) ([]*model.Technology, error) {
	tech := &model.Technology{
		ID:   fmt.Sprintf("T%d", rand.Intn(100)),
		Text: "Testing goodness",
	}
	r.technologies = append(r.technologies, tech)
	return []*model.Technology{tech}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
