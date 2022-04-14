package model

import (
	"api/internal/application/command"
)

type PC struct {
	Model string
}

func NewPC(cmd command.PCCommand) *PC {
	return &PC{cmd.Model}
}
