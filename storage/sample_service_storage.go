package storage

import "test/protos"

type PersonStorage interface {
	ListPeople() ([]*protos.Person, error)
}
