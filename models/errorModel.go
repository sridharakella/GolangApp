package models

type ErrorModel struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	//Message string `json:"message"`
}
