package wish_list_link

import "context"

func (wllr *WishListLinksRepo) DeleteWishListLink(ctx context.Context, link string) error {
	const query = "DELETE FROM wish_list_links WHERE link = $1"

	_, err := wllr.db.Exec(ctx, query, link)
	if err != nil {
		return err
	}

	return nil
}
