package model

import (
  "api/internal/application/command"
)

type Person struct {
  Id    PersonId
  Name  PersonName
  Email PersonEmail
  PC    *PC
  Sake  []Sake
  State PersonState
}

func NewPerson(cmd command.PersonCommand) (*Person) {
  var sakes []Sake
  for _, v := range cmd.Sake {
    sake := NewSake(v)
    sakes = append(sakes, *sake)
  }
  return &Person{
    Id: NewPersonId(),
    Name: NewPersonName(cmd.PersonName),
    Email: NewPersonEmail(cmd.PersonEmail),
    PC: NewPC(cmd.PC),
    Sake: sakes,
    State: NewPersonState("Alive"),
  }
}
