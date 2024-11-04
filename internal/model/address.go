package model

type Address struct {
	Id        uint32
	ProfileId uint32
	City      string
	Street    string
	House     string
	Flat      *string
}
