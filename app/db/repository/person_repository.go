package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"person-service/config"
	"person-service/db/entity"
	"person-service/utils"
	"strconv"
	"time"
)

type PersonRepositoryImpl struct {
	db *sql.DB
}

const TimeFormat = "2006-01-02 15:04:05.000000000"

func New(datasource config.Datasource) (*PersonRepositoryImpl, error) {
	const op = "storage.postgres.New"

	connection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		datasource.Host,
		datasource.Port,
		datasource.User,
		datasource.Password,
		datasource.DbName,
	)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS person(
	    	id 			uuid 		PRIMARY KEY,
	    	first_name 	text 		NOT NULL,
	    	last_name 	text 		NOT NULL,
	    	age 		int			NOT NULL,
	    	last_update timestamp 	NOT NULL 
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &PersonRepositoryImpl{db: db}, nil
}

// DeletePerson delete person with selected id.
func (s *PersonRepositoryImpl) DeletePerson(id uuid.UUID) (string, error) {
	const op = "storage.postgres.DeletePerson"

	sqlStatement := `DELETE FROM person WHERE id = $1`
	_, err := s.db.Exec(sqlStatement, id.String())
	if err != nil {
		return "", fmt.Errorf("error while delete person: %s: %w", op, err)
	}

	return id.String(), nil
}

// FindPersonById find person by id.
func (s *PersonRepositoryImpl) FindPersonById(id *uuid.UUID) (entity.Person, error) {
	const op = "storage.postgres.FindPersonById"

	var person entity.Person

	sqlStatement := `SELECT * FROM person p WHERE p.id = $1`
	err := s.db.QueryRow(sqlStatement, id.String()).
		Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age, &person.Timestamp, &person.Login)

	if err != nil {
		return entity.Person{}, fmt.Errorf("error while find person: %s: %w", op, err)
	} else {
		return person, nil
	}
}

// FindPersonByLogin find person by login.
func (s *PersonRepositoryImpl) FindPersonByLogin(login string) (entity.Person, error) {
	const op = "storage.postgres.FindPersonByLogin"

	var person entity.Person

	sqlStatement := `SELECT * FROM person p WHERE p.login = $1`
	err := s.db.QueryRow(sqlStatement, login).
		Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age, &person.Timestamp, &person.Login)

	if err != nil {
		return entity.Person{}, fmt.Errorf("error while find person: %s: %w", op, err)
	} else {
		return person, nil
	}
}

// UpdatePerson method update existing person in database or creates new if id of argument is null.
func (s *PersonRepositoryImpl) UpdatePerson(person entity.Person) (entity.Person, error) {
	const op = "storage.postgres.UpdatePerson"

	if person.Id == nil {
		return s.SavePerson(person)
	}

	var updatedPerson entity.Person
	sqlStatement := `UPDATE person p SET first_name = $1, last_name = $2, age=$3, last_update = $4 
              WHERE id = $5 
              RETURNING p.id, p.first_name, p.last_name, p.age, p.last_update`
	timestamp := time.Now().Format(TimeFormat)

	err := s.db.QueryRow(sqlStatement, person.FirstName, person.LastName, person.Age, timestamp, person.Id).
		Scan(&updatedPerson.Id, &updatedPerson.FirstName, &updatedPerson.LastName, &updatedPerson.Age, &updatedPerson.Timestamp)

	if err != nil {
		return entity.Person{}, fmt.Errorf("error while update existing person: %s: %w", op, err)
	} else {
		return updatedPerson, nil
	}
}

// SavePerson save new person to database or updated existing row.
func (s *PersonRepositoryImpl) SavePerson(p entity.Person) (entity.Person, error) {
	const op = "storage.postgres.SavePerson"
	var person entity.Person

	var id string
	sqlStatement := `INSERT INTO person(id, first_name, last_name, age, last_update) 
						VALUES ($1, $2, $3, $4, $5) 
							RETURNING id, first_name, last_name, age, last_update`
	timestamp := time.Now().Format(TimeFormat)

	if utils.IsNullableUUID(p.Id) {
		id = uuid.New().String()
	} else {
		id = p.Id.String()
	}

	err := s.db.QueryRow(sqlStatement, id, p.FirstName, p.LastName, p.Age, timestamp).
		Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age, &person.Timestamp)

	if err != nil {
		return entity.Person{}, fmt.Errorf("error while save new person: %s: %w", op, err)
	} else {
		return person, nil
	}
}

// LoadPersons load first 50 persons from database.
func (s *PersonRepositoryImpl) LoadPersons(page *string) ([]entity.Person, error) {
	const op = "storage.postgres.LoadPersons"

	var person entity.Person
	pageInt, _ := strconv.Atoi(*page)
	sqlStatement := `SELECT p.id, p.first_name, p.last_name, p.age, p.last_update, p.login FROM person p LIMIT 50 OFFSET $1`

	var offset int
	if pageInt <= 1 {
		offset = 0
	} else {
		offset = (pageInt - 1) * 50
	}

	rows, err := s.db.Query(sqlStatement, offset)
	if err != nil {
		return nil, fmt.Errorf("error whole load persons: %s: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var persons []entity.Person
	for rows.Next() {
		err := rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age, &person.Timestamp, &person.Login)
		if err != nil {
			log.Fatal(err)
		}

		persons = append(persons, entity.Person{
			Id:        person.Id,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Age:       person.Age,
			Timestamp: person.Timestamp,
			Login:     person.Login,
		})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return persons, nil
}
