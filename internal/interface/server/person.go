package server

import (
	"context"
	"errors"
	"fmt"

	"api/internal/application/service"
	"api/proto/person/protobuf"
)

type personManagementServer struct {
	personAS service.PersonApplicationServiceIF
}

func NewPersonManagementServer(personAS service.PersonApplicationServiceIF) protobuf.PersonManagementServer {
	return &personManagementServer{personAS}
}

func (server *personManagementServer) Create(ctx context.Context, req *protobuf.CreatePersonRequest) (*protobuf.PersonBaseResponse, error) {
	fmt.Printf("value: %v, type: %T", req, req)

	// server.personAS.Create(ctx, cmd)
	return nil, errors.New("method Create not implemented")
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
