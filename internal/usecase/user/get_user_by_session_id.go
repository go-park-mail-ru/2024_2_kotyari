package user

import "context"

func (us *UsersService) GetUserBySessionID(ctx context.Context, sessionID string) (string, string, error) {
	session, err := us.sessionService.SessionRepo.Get(ctx, sessionID)
	if err != nil {
		return "", "", err
	}

	user, err := us.userRepo.GetUserByUserID(ctx, session.UserID)
	if err != nil {
		return "", "", err
	}

	return user.Username, user.City, nil
}
