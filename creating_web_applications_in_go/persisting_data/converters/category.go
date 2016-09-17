package converters

import (
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/models"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/viewmodels"
)

func ConvertCategoyToViewModel(category models.Category) viewmodels.Category {
	result := viewmodels.Category{
		ImageUrl:      category.ImageUrl(),
		Title:         category.Title(),
		Description:   category.Description(),
		IsOrientRight: category.IsOrientRight(),
		Id:            category.Id(),
	}

	for _, p := range category.Products() {
		result.Products = append(result.Products, ConvertProductToViewModel(p))
	}

	return result
}
