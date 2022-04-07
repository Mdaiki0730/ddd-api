package model

type PersonState string

func NewPersonState(state string) PersonState {
  return PersonState(state)
}
