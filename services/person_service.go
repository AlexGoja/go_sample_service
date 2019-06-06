package services

import (
	"context"
	"net/http"
	"test/protos"
	"test/storage"
)

type PersonService struct {
	storage storage.PersonStorage
}

func NewPersonService(personStorage storage.PersonStorage) *PersonService {
	return &PersonService{storage:personStorage}
}

func (pService *PersonService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (pService *PersonService) ListPeople(ctx context.Context, in *protos.ListPersonRequest) (*protos.ListPersonResponse, error) {

	var people []*protos.Person
	var err error
	if people, err = pService.storage.ListPeople(); err != nil {
		return nil, err
	}

	return &protos.ListPersonResponse{Person:people}, nil
}
