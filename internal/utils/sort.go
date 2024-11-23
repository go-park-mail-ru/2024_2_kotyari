package utils

const (
	SearchFieldParam = "sort"
	SearchOrderParam = "order"

	ascSortOrderOption  = "asc"
	descSortOrderOption = "desc"
)

func ReturnSortOrderOption(sortOrderOption string) string {
	if sortOrderOption != ascSortOrderOption && sortOrderOption != descSortOrderOption {
		return descSortOrderOption
	}

	return sortOrderOption
}
