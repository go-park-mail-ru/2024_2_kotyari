package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/address"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type ProfileResponse struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Age       uint8  `json:"age"`
	AvatarUrl string `json:"avatar_url"`
	Address   address.AddressResponse
}

func FromModel(profile model.Profile) ProfileResponse {

	return ProfileResponse{
		ID:       profile.ID,
		Email:    profile.Email,
		Username: profile.Username,
		Age:      profile.Age,
		Address: address.AddressResponse{
			ID:     profile.Address.Id,
			City:   profile.Address.City,
			Street: profile.Address.Street,
			House:  profile.Address.House,
			Flat:   *profile.Address.Flat,
		},
		Gender:    profile.Gender,
		AvatarUrl: profile.AvatarURL,
	}
}

type UpdateProfileRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

type AvatarResponse struct {
	AvatarUrl string `json:"avatar_url"`
}
