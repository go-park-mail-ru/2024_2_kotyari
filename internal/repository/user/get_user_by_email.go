package user

import "context"

func (ur *UsersRepo) GetUserByEmail(ctx context.Context, email string) (string, error) {
	const query = `
	select username from users where users.email = ? limit 1 	
	`

	var username string

	err := ur.db.QueryRow(ctx, query, email).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}
