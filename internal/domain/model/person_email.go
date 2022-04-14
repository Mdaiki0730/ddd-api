package model

type PersonEmail string

func NewPersonEmail(email string) PersonEmail {
	return PersonEmail(email)
}
