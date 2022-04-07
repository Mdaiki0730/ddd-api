package model

import (
  "api/pkg/randuuid"
)

type PersonId string

func NewPersonId() PersonId {
  id := randuuid.New()
  return PersonId(id)
}
