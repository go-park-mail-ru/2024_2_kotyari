package wish_list

type addProductToWishlistsRequest struct {
	ProductId uint32   `json:"product_id"`
	Links     []string `json:"links"`
}

type copyWishlistsRequest struct {
	Link string `json:"link"`
}

type createWishlistRequest struct {
	Name string `json:"name"`
}

type deleteWishlistsRequest struct {
	Link string `json:"link"`
}

type getWishlistByLinkResponse struct {
	IsAuthor bool   `json:"is_creator"`
	Link     string `json:"link"`
	Name     string `json:"name"`
	Items    []item `json:"items"`
}

type item struct {
	Id       uint32 `json:"id"`
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
	Price    uint32 `json:"price"`
}

type removeFromWishlistRequest struct {
	Links     []string `json:"links"`
	ProductId uint32   `json:"product_id"`
}

type renameWishlistRequest struct {
	NewName string `json:"new_name"`
	Link    string `json:"link"`
}
