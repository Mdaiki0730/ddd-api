package repoif

import (
	"api/internal/domain/model"
	"context"
)

type Person interface {
	InsertOne(ctx context.Context, person *model.Person) error
	UpdateOne(ctx context.Context, person *model.Person) (*model.Person, error)
	DeleteById(ctx context.Context, personId model.PersonId) error
	FindById(ctx context.Context, personId model.PersonId) (*model.Person, error)
	FindAll(ctx context.Context) ([]model.Person, error)
}
