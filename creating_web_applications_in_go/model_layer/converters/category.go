package converters

import (
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/model_layer/models"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/model_layer/viewmodels"
)

func ConvertCategoryToViewModel(category models.Category) viewmodels.Category {
	result := viewmodels.Category{
		ImageUrl:      category.ImageUrl(),
		Title:         category.Title(),
		Description:   category.Description(),
		IsOrientRight: category.IsOrientRight(),
		Id:            category.Id(),
	}

	return result
}
