package csat

import grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"

type CreateCsatRequest struct {
	Type   string `json:"type"`
	Rating uint32 `json:"rating"`
	Text   string `json:"text"`
}

type GetCsatsRequest struct {
	Type   string `json:"type"`
	Rating uint32 `json:"rating"`
	Text   string `json:"text"`
}

type GetCsatsResponse struct {
	Text string `json:"text"`
}

type GetStatisticsRequest struct {
	Type string `json:"rating"`
}

type CreateCsatResponse struct {
	Type   string `json:"type"`
	Rating uint32 `json:"rating"`
	Text   string `json:"text"`
}

type StarVotes struct {
	StarNumber uint32 `json:"star_number"`
	VoteCount  uint32 `json:"vote_count"`
}

type GetStatisticsResponse struct {
	Stats   []*StarVotes `json:"stats"`
	Average float64      `json:"average"`
}

func (cr *CreateCsatRequest) ToGrpcCreateCsatRequest() *grpc_gen.CreateCsatRequest {
	return &grpc_gen.CreateCsatRequest{
		Type:   cr.Type,
		Rating: cr.Rating,
		Text:   cr.Text,
	}
}

func FromGrpcResponse(gr *grpc_gen.CreateCsatResponse) CreateCsatResponse {
	return CreateCsatResponse{
		Type:   gr.GetType(),
		Rating: gr.GetRating(),
		Text:   gr.GetText(),
	}
}

//func (cr *CreateCsatRequest) ToModel() model.User {
//	return model.User{
//		Email:    cr.Email,
//		Username: cr.Username,
//		Password: cr.Password,
//	}
//}
