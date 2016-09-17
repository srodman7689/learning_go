package converters

import (
	"testing"

	"github.com/srodman7689/learning_go/creating_web_applications_in_go/model_layer/models"
)

func Test_ConvertsCategoryToViewModel(t *testing.T) {
	category := models.Category{}
	category.SetImageUrl("the image URL")
	category.SetTitle("the title")
	category.SetDescription("the description")
	category.SetIsOrientRight(true)
	category.SetId(42)

	result := ConvertCategoryToViewModel(category)

	if result.ImageUrl != category.ImageUrl() {
		t.Log("Image URL not converted properly")
		t.Fail()
	}

	if result.Title != category.Title() {
		t.Log("Title not converted properly")
		t.Fail()
	}

	if result.Description != category.Description() {
		t.Log("Description not converted properly")
		t.Fail()
	}

	if result.IsOrientRight != category.IsOrientRight() {
		t.Log("Is Orient Right not converted properly")
		t.Fail()
	}

	if result.Id != category.Id() {
		t.Log("Id not converted properly")
		t.Fail()
	}
}
