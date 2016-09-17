package converters

import (
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/models"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/viewmodels"
)

func ConvertProductToViewModel(product models.Product) viewmodels.Product {
	result := viewmodels.Product{
		Name:             product.Name(),
		DescriptionShort: product.DescriptionShort(),
		DescriptionLong:  product.DescriptionLong(),
		PricePerLiter:    product.PricePerLiter(),
		PricePer10Liter:  product.PricePer10Liter(),
		Origin:           product.Origin(),
		IsOrganic:        product.IsOrganic(),
		ImageUrl:         product.ImageUrl(),
		Id:               product.Id(),
	}

	return result
}
