package models

type Card struct {
	CardNumber     string `validate:"required"`
	CustomerName   string `validate:"required"`
	Holder         string `validate:"required"`
	ExpirationDate string `validate:"required"`
	SecurityCode   string `validate:"required"`
	Brand          string `validate:"required"`
}
