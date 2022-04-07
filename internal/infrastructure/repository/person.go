package repository

import (
	"context"

	"api/internal/domain/model"
	"api/internal/domain/repoif"
)

type person struct {}

// use only dependency
func NewPerson() repoif.Person {
	return &person{}
}

func (r person) InsertOne(ctx context.Context, person *model.Person) error {
	// write code
	return nil
}

func (r person) UpdateOne(ctx context.Context, person *model.Person) (*model.Person, error) {
	// write code
	return &model.Person{}, nil
}

func (r person) DeleteById(ctx context.Context, personId model.PersonId) error {
	// write code
	return nil
}

func (r person) FindById(ctx context.Context, personId model.PersonId) (*model.Person, error) {
	// write code
	return &model.Person{}, nil
}

func (r person) FindAll(ctx context.Context) ([]model.Person, error) {
	// write code
	return []model.Person{}, nil
}
