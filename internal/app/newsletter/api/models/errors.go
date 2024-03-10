package models

type BaseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
