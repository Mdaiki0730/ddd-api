package model

import (
	"api/internal/application/command"
)

type Person struct {
	Id    PersonId    `json:"person_id"`
	Name  PersonName  `json:"person_name"`
	Email PersonEmail `json:"person_email"`
	PC    *PC         `json:"pc"`
	Sake  []Sake      `json:"sake"`
	State PersonState `json:"state"`
}

func NewPerson(cmd command.PersonCommand) *Person {
	var sakes []Sake
	for _, v := range cmd.Sake {
		sake := NewSake(v)
		sakes = append(sakes, *sake)
	}
	return &Person{
		Id:    NewPersonId(),
		Name:  NewPersonName(cmd.PersonName),
		Email: NewPersonEmail(cmd.PersonEmail),
		PC:    NewPC(cmd.PC),
		Sake:  sakes,
		State: NewPersonState("Alive"),
	}
}
