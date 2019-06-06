package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"test/protos"
)

type PersonStorage struct {
	conn *sql.DB
}

func NewPersonStorage(postgreStr string) *PersonStorage {
	conn, err := sql.Open("postgres", postgreStr)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	return &PersonStorage{conn}
}

func (p *PersonStorage) ListPeople() ([]*protos.Person, error) {
	const stmt = "select * from users"

	trans, err := p.conn.Begin()

	if err != nil {
		return nil, err
	}

	rows, err := trans.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer trans.Commit()

	var people []*protos.Person

	for rows.Next() {
		var person protos.Person
		if err := rows.Scan(&person.Id, &person.Name, &person.Address, &person.Username, &person.Password); err !=nil {
			return nil, err
		}
		people = append(people, &person)
	}

	fmt.Println(people)

	return people, nil
}
