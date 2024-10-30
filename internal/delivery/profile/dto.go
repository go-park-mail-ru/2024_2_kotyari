package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/address"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type ProfileDTO struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Age       uint8  `json:"age"`
	AvatarUrl string `json:"avatar_url"`
}

type ProfileRequest struct {
	ID uint32 `json:"id"`
}

type ProfileResponse struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Age       uint8  `json:"age"`
	AvatarUrl string `json:"avatar_url"`
	Address   address.AddressDTO
}

func (p *ProfileDTO) ToModel() model.Profile {
	return model.Profile{
		ID:        p.ID,
		Email:     p.Email,
		Username:  p.Username,
		Gender:    p.Gender,
		Age:       p.Age,
		AvatarUrl: p.AvatarUrl,
	}
}

func FromModel(profile model.Profile) ProfileDTO {
	return ProfileDTO{
		ID:        profile.ID,
		Email:     profile.Email,
		Username:  profile.Username,
		Gender:    profile.Gender,
		Age:       profile.Age,
		AvatarUrl: profile.AvatarUrl,
	}
}

type UpdateProfileRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

func (r *UpdateProfileRequest) ToModel(profileID uint32) model.Profile {
	return model.Profile{
		ID:       profileID,
		Email:    r.Email,
		Username: r.Username,
		Gender:   r.Gender,
	}
}
