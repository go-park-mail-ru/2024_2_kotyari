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

func optionToDTO(option model.Option) dtoOption {
	return dtoOption{
		Link:  option.Link,
		Value: option.Value,
	}
}

func optionsBlockToDTO(optionsBlock model.OptionsBlock) dtoOptionBlock {
	options := make([]dtoOption, len(optionsBlock.Options))
	for i, option := range optionsBlock.Options {
		options[i] = optionToDTO(option)
	}
	return dtoOptionBlock{
		Title:   optionsBlock.Title,
		Type:    optionsBlock.Type,
		Options: options,
	}
}

func optionsToDto(opts model.Options) []dtoOptionBlock {
	res := make([]dtoOptionBlock, len(opts.Values))

	for i, option := range opts.Values {
		res[i] = optionsBlockToDTO(option)
	}

	return res
}
