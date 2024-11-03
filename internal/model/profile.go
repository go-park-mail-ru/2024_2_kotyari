package model

type Profile struct {
	ID        uint32
	Email     string
	Username  string
	Gender    string
	Address   AddressDTO
	Age       uint8
	AvatarURL string
}
