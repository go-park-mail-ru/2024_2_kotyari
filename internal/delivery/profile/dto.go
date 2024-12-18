package profile

import (
	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/address"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type UsersDefaultResponse struct {
	Username  string `json:"username"`
	City      string `json:"city"`
	AvatarUrl string `json:"avatar_url"`
}

type ProfilesResponse struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	AvatarUrl string `json:"avatar_url"`
	Address   address.AddressResponse
}

func fromGrpcResponse(p *profilegrpc.GetProfileResponse, addr model.Address) ProfilesResponse {
	return ProfilesResponse{
		ID:        p.UserId,
		Email:     p.Email,
		Username:  p.Username,
		Gender:    p.Gender,
		AvatarUrl: p.AvatarUrl,
		Address: address.AddressResponse{
			Address: addr.Text,
		},
	}
}

//easyjson:json
type UpdateProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

type AvatarResponse struct {
	AvatarUrl string `json:"avatar_url"`
}

//easyjson:json
type UpdatePasswordRequest struct {
	OldPassword    string `json:"old_password"`
	NewPassword    string `json:"new_password"`
	RepeatPassword string `json:"repeat_password"`
}
