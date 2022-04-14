package server

import (
	"context"
	"errors"

	"api/internal/application/service"
  "api/internal/application/command"
  "api/pkg/jsonutil"
	"api/proto/person/protobuf"
)

type personManagementServer struct {
	personAS service.PersonApplicationServiceIF
}

func NewPersonManagementServer(personAS service.PersonApplicationServiceIF) protobuf.PersonManagementServer {
	return &personManagementServer{personAS}
}

func (server *personManagementServer) Create(ctx context.Context, req *protobuf.CreatePersonRequest) (*protobuf.PersonBaseResponse, error) {
  cmd := command.PersonCommand{}
  jsonutil.DataTransfer(req, &cmd)

	result, err := server.personAS.Create(ctx, cmd)
  if err != nil {
    return nil, err
  }

  res := &protobuf.PersonBaseResponse{}
  jsonutil.DataTransfer(result, res)
	return res, nil
}
func (server *personManagementServer) Get(context.Context, *protobuf.GetPersonRequest) (*protobuf.PersonBaseResponse, error) {
	return nil, errors.New("method Get not implemented")
}
func (server *personManagementServer) List(context.Context, *protobuf.ListPersonRequest) (*protobuf.ListPersonResponse, error) {
	return nil, errors.New("method List not implemented")
}
func (server *personManagementServer) Update(context.Context, *protobuf.UpdatePersonRequest) (*protobuf.PersonBaseResponse, error) {
	return nil, errors.New("method Update not implemented")
}
func (server *personManagementServer) Delete(context.Context, *protobuf.DeletePersonRequest) (*protobuf.DeletePersonResponse, error) {
	return nil, errors.New("method Delete not implemented")
}
