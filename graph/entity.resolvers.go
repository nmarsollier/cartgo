package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"

	"github.com/nmarsollier/cartgo/graph/model"
	"github.com/nmarsollier/cartgo/graph/resolvers"
)

// FindCartByID is the resolver for the findCartByID field.
func (r *entityResolver) FindCartByID(ctx context.Context, id string) (*model.Cart, error) {
	cart, err := resolvers.FindCartById(ctx, id)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// Entity returns model.EntityResolver implementation.
func (r *Resolver) Entity() model.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }