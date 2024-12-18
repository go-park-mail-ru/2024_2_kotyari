package wish_list

//easyjson:json
type addProductToWishlistsRequest struct {
	ProductId uint32   `json:"product_id"`
	Links     []string `json:"links"`
}

//easyjson:json
type copyWishlistsRequest struct {
	Link string `json:"link"`
}

//easyjson:json
type createWishlistRequest struct {
	Name string `json:"name"`
}

//easyjson:json
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

//easyjson:json
type removeFromWishlistRequest struct {
	Links     []string `json:"links"`
	ProductId uint32   `json:"product_id"`
}

//easyjson:json
type renameWishlistRequest struct {
	NewName string `json:"new_name"`
	Link    string `json:"link"`
}
