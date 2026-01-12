package fromdb

import (
	_ "embed"
)

//go:embed build_data/get_categories.json
var getCategoriesJSON []byte

func GetCategories() []byte{
	return getCategoriesJSON
}