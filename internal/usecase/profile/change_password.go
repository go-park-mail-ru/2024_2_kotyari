package profile

import "context"

// todo

func (ps *ProfilesService) ChangePassword(ctx context.Context, userId uint32, oldPassword, newPassword, repeatPassword string) error {
	user, err := ps.userGetter.GetUserByUserID(ctx, userId)
	_ = user

	err = ps.profileRepo.ChangePassword(ctx, userId, newPassword)
	if err != nil {
		return err
	}

	return nil
}
