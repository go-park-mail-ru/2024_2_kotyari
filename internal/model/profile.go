package model

type Profile struct {
	ID        uint32
	Email     string
	Username  string
	Gender    string
	Address   Address
	Age       uint8
	AvatarUrl string
}
