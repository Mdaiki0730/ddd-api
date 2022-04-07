package command

type PersonCommand struct {
  PersonName  string `json:"person_name"`
  PersonEmail string `json:"person_email"`
  PC          PCCommand `json:"pc"`
  Sake        []SakeCommand `json:"sake"`
}

type PCCommand struct {
  Model string `json:"model"`
}

type SakeCommand struct {
  Name    string `json:"name"`
  Percent string `json:"percent"`
  Liquor  string `json:"liquor"`
}
