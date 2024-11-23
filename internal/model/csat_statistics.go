package model

type CSATStatistics struct {
	Ratings []CSATRating
	Total   uint32
	Avg     float64
}

type CSATRating struct {
	RatingName  uint32 // 1-10
	RatingValue uint32 //количество
}
