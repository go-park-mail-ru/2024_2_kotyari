package wish_list_link

import "context"

func (wllr *WishListLinksRepo) CreateLink(ctx context.Context, userID uint32, link string) error {
	_, err := wllr.db.Exec(ctx, "INSERT INTO wish_list_links (user_id,link) VALUES ($1, $2)", userID, link)
	if err != nil {
		return err
	}

	return nil
}
