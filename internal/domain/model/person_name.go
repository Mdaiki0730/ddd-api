package model

type PersonName string

func NewPersonName(name string) PersonName {
  return PersonName(name) 
}
