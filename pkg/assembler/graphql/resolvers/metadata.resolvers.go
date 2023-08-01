package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestHasMetadata is the resolver for the ingestHasMetadata field.
func (r *mutationResolver) IngestHasMetadata(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType *model.MatchFlags, hasMetadata model.HasMetadataInputSpec) (*model.HasMetadata, error) {
	return r.Backend.IngestHasMetadata(ctx, subject, pkgMatchType, hasMetadata)
}

// HasMetadata is the resolver for the HasMetadata field.
func (r *queryResolver) HasMetadata(ctx context.Context, hasMetadataSpec model.HasMetadataSpec) ([]*model.HasMetadata, error) {
	return r.Backend.HasMetadata(ctx, &hasMetadataSpec)
}