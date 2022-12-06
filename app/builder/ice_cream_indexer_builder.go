package builder

import (
	"ic-indexer-service/models"
	"ic-service/app/model/bo"
	"strings"
)

type IcecreamIndexerBuilderInterface interface {
	IcecreamIndexerBuilder(*bo.Icecream, bool) models.IcecreamClientRequest
}

type icecreamIndexerBuilder struct {
}

func NewIcecreamIndexerBuilder() IcecreamIndexerBuilderInterface {
	return &icecreamIndexerBuilder{}
}

func (iib icecreamIndexerBuilder) IcecreamIndexerBuilder(request *bo.Icecream, isDeleted bool) models.IcecreamClientRequest {
	// new icecream builder

	clientRequest := models.IcecreamClientRequest{}

	clientRequest.Description = request.Description
	clientRequest.Name = request.Name
	clientRequest.AllergyInfo = request.AllergyInfo
	clientRequest.ImageClosed = request.ImageClosed
	clientRequest.ImageOpen = request.ImageOpen

	if len(request.Ingredients) > 0 {
		clientRequest.Ingredients = strings.Split(request.Ingredients, ",")
	}

	if len(request.SourcingValues) > 0 {
		clientRequest.SourcingValues = strings.Split(request.SourcingValues, ",")
	}

	clientRequest.Story = request.Story
	clientRequest.DietaryCertifications = request.DietaryCertifications
	clientRequest.Description = request.Description
	clientRequest.ID = request.Id
	clientRequest.ProductID = request.ProductId
	clientRequest.IsDeleted = isDeleted
	return clientRequest
}
