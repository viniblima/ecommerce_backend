package models

type PaymentMethod int

const (
	CreditCard = iota
	Ticket
	Pix
)
