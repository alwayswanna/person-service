package mappers

import (
	"person-service/db/entity"
	"person-service/model"
)

func ToPerson(request model.PersonRequest) entity.Person {
	return entity.Person{
		Id:        &request.Id,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Age:       request.Age,
		Timestamp: &request.Timestamp,
	}
}

func ToPersonResponse(entity entity.Person) model.PersonResponse {
	return model.PersonResponse{
		Id:        *entity.Id,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Age:       entity.Age,
		Timestamp: *entity.Timestamp,
		Login:     entity.Login,
	}
}

func ToPersonsResponse(entities []entity.Person) []model.PersonResponse {
	persons := make([]model.PersonResponse, len(entities))
	for index, person := range entities {
		persons[index] = ToPersonResponse(person)
	}
	return persons
}
