package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql-mongodb-demo/graph/generated"
	"go-graphql-mongodb-demo/graph/model"
	"go-graphql-mongodb-demo/repository"
	"time"
)

var userRepo repository.UserRepository = repository.Init()

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: time.Now().String(),
	}

	userRepo.Save(user)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return userRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
