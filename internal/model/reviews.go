package model

type Reviews struct {
	TotalReviewCount uint32
	TotalRating      float32
	UserReview       Review
	Reviews          []Review
}
