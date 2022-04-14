package dto

type PersonBase struct {
	PersonName  string     `json:"person_name"`
	PersonEmail string     `json:"person_email"`
	Pc          PCBase     `json:"pc"`
	Sake        []SakeBase `json:"sake"`
}

type PCBase struct {
	Model string `json:"model"`
}

type SakeBase struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
	Liquor  string `json:"liquor"`
}
