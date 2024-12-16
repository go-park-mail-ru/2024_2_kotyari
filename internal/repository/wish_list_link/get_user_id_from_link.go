package wish_list_link

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (wllr *WishListLinksRepo) GetUserIDFromLink(ctx context.Context, link string) (uint32, error) {
	const query = "SELECT user_id FROM wish_list_links WHERE link = $1"

	var userID uint32

	err := wllr.db.QueryRow(ctx, query, link).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("wishlist not found")
		}
		return 0, fmt.Errorf("failed to query PostgreSQL: %w", err)
	}

	return userID, nil
}
