package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestBuilder is the resolver for the ingestBuilder field.
func (r *mutationResolver) IngestBuilder(ctx context.Context, builder *model.IDorBuilderInput) (string, error) {
	return r.Backend.IngestBuilder(ctx, builder)
}

// IngestBuilders is the resolver for the ingestBuilders field.
func (r *mutationResolver) IngestBuilders(ctx context.Context, builders []*model.IDorBuilderInput) ([]string, error) {
	return r.Backend.IngestBuilders(ctx, builders)
}

// Builders is the resolver for the builders field.
func (r *queryResolver) Builders(ctx context.Context, builderSpec model.BuilderSpec) ([]*model.Builder, error) {
	return r.Backend.Builders(ctx, &builderSpec)
}

// BuildersList is the resolver for the buildersList field.
func (r *queryResolver) BuildersList(ctx context.Context, builderSpec model.BuilderSpec, after *string, first *int) (*model.BuilderConnection, error) {
	return r.Backend.BuildersList(ctx, builderSpec, after, first)
}
