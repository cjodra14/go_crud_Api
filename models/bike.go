package models

type Bike struct {
	Model        string `json:"model"`
	Displacement string `json:"displacement"`
	Brand        *Brand `json:"brand"`
}

type Brand struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
