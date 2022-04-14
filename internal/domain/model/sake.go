package model

import (
	"api/internal/application/command"
)

type Sake struct {
	Name    string
	Percent string
	Liquor  string
}

func NewSake(cmd command.SakeCommand) *Sake {
	return &Sake{
		Name:    cmd.Name,
		Percent: cmd.Percent,
		Liquor:  cmd.Liquor,
	}
}
