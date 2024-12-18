package product

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type dtoOption struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

func (dtoOption *dtoOption) ToModel() model.Option {
	return model.Option{
		Link:  dtoOption.Link,
		Value: dtoOption.Value,
	}
}

type dtoOptionBlock struct {
	Title   string      `json:"title"`
	Type    string      `json:"type"`
	Options []dtoOption `json:"options"`
}

func (dto *dtoOptionBlock) ToModel() model.OptionsBlock {
	options := make([]model.Option, len(dto.Options))
	for i, opt := range dto.Options {
		options[i] = opt.ToModel()
	}
	return model.OptionsBlock{
		Title:   dto.Title,
		Type:    dto.Type,
		Options: options,
	}
}
