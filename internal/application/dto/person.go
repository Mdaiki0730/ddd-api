package dto

type PersonBase struct {
  PersonId    string     `json:"person_id"`
	PersonName  string     `json:"person_name"`
	PersonEmail string     `json:"person_email"`
	Pc          PCBase     `json:"pc"`
	Sake        []SakeBase `json:"sake"`
  State       string     `json:"state"`
}

type PCBase struct {
	Model string `json:"model"`
}

type SakeBase struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
	Liquor  string `json:"liquor"`
}
