package builder

import (
	"ic-service/app/helpers"
	"ic-service/app/model/bo"
	"ic-service/app/model/request"
	"sort"
	"strings"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=IcecreamBuilderInterface"
type IcecreamBuilderInterface interface {
	IcecreamBuilder(request.IcecreamRequest) *bo.Icecream
	IcecreamPartialBuilder(request.IcecreamUpdateRequest, *bo.Icecream) *bo.Icecream
}

type icecreamBuilder struct {
}

func NewIcecreamBuilder() IcecreamBuilderInterface {
	return &icecreamBuilder{}
}

func (alb icecreamBuilder) IcecreamBuilder(request request.IcecreamRequest) *bo.Icecream {
	// new icecream builder
	newIcecream := bo.NewIcecream()
	newIcecream.ProductId = request.ProductId
	newIcecream.Name = request.Name
	newIcecream.AllergyInfo = request.AllergyInfo
	newIcecream.ImageClosed = request.ImageClosed
	newIcecream.ImageOpen = request.ImageOpen

	if request.Ingredients != nil && len(request.Ingredients) > 0 {
		sort.Strings(request.Ingredients)
		newIcecream.Ingredients = strings.Join(request.Ingredients, ",")
	}

	if request.SourcingValues != nil && len(request.SourcingValues) > 0 {
		sort.Strings(request.SourcingValues)
		newIcecream.SourcingValues = strings.Join(request.SourcingValues, ",")
	}

	newIcecream.Story = request.Story
	newIcecream.DietaryCertifications = request.DietaryCertifications
	newIcecream.Description = request.Description
	newIcecream.Id = request.Id

	return newIcecream
}

func (alb icecreamBuilder) IcecreamPartialBuilder(icecreamUpdateRequest request.IcecreamUpdateRequest, oldIcecream *bo.Icecream) *bo.Icecream {
	// new icecream update builder
	var empty []string
	oldIcecream.ProductId = helpers.GetString(icecreamUpdateRequest.ProductId, oldIcecream.ProductId)
	oldIcecream.Name = helpers.GetString(icecreamUpdateRequest.Name, oldIcecream.Name)
	oldIcecream.AllergyInfo = helpers.GetString(icecreamUpdateRequest.AllergyInfo, oldIcecream.AllergyInfo)
	oldIcecream.ImageClosed = helpers.GetString(icecreamUpdateRequest.ImageClosed, oldIcecream.ImageClosed)
	oldIcecream.ImageOpen = helpers.GetString(icecreamUpdateRequest.ImageOpen, oldIcecream.ImageOpen)

	sourcingValues := helpers.GetArrayString(icecreamUpdateRequest.SourcingValues, empty)
	sourcingValuesString := strings.Join(sourcingValues, string(','))
	if len(sourcingValues) > 0 {
		oldIcecream.Ingredients = helpers.GetString(request.CusString{Value: &sourcingValuesString, Set: true}, oldIcecream.SourcingValues)
	} else {
		oldIcecream.Ingredients = helpers.GetString(request.CusString{Value: &sourcingValuesString, Set: false}, oldIcecream.SourcingValues)
	}

	ingredients := helpers.GetArrayString(icecreamUpdateRequest.Ingredients, empty)
	ingredientsString := strings.Join(ingredients, string(','))
	if len(ingredients) > 0 {
		oldIcecream.Ingredients = helpers.GetString(request.CusString{Value: &ingredientsString, Set: true}, oldIcecream.Ingredients)
	} else {
		oldIcecream.Ingredients = helpers.GetString(request.CusString{Value: &ingredientsString, Set: false}, oldIcecream.Ingredients)
	}
	helpers.GetString(icecreamUpdateRequest.Story, oldIcecream.Story)
	oldIcecream.DietaryCertifications = helpers.GetString(icecreamUpdateRequest.DietaryCertifications, oldIcecream.DietaryCertifications)
	oldIcecream.Description = helpers.GetString(icecreamUpdateRequest.Description, oldIcecream.Description)

	return oldIcecream
}
