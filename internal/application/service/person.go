package service

import (
	"context"

	"api/internal/application/command"
	"api/internal/application/dto"
  "api/internal/domain/model"
	"api/internal/domain/repoif"
  "api/pkg/jsonutil"
)

type PersonApplicationServiceIF interface {
	Create(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error)
	Get(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error)
	List(ctx context.Context) ([]dto.PersonBase, error)
	Update(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error)
	Delete(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error)
}

type PersonApplicationService struct {
	personRepository repoif.Person
}

func NewPersonApplicationService(r repoif.Person) PersonApplicationServiceIF {
	return &PersonApplicationService{r}
}

func (personAS *PersonApplicationService) Create(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error) {
  // create person instance
  person := model.NewPerson(cmd)

  if err := personAS.personRepository.InsertOne(ctx, person); err != nil {
    return nil, err
  }

  res := &dto.PersonBase{}
  jsonutil.DataTransfer(person, res)
	return res, nil
}

func (personAS *PersonApplicationService) Get(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error) {
	return &dto.PersonBase{}, nil
}

func (personAS *PersonApplicationService) List(ctx context.Context) ([]dto.PersonBase, error) {
	return []dto.PersonBase{}, nil
}

func (personAS *PersonApplicationService) Update(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error) {
	return &dto.PersonBase{}, nil
}

func (personAS *PersonApplicationService) Delete(ctx context.Context, cmd command.PersonCommand) (*dto.PersonBase, error) {
	return &dto.PersonBase{}, nil
}
