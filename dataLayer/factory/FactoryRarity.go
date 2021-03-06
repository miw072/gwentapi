package factory

import (
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/models"
	"github.com/tri125/gwentapi/helpers"
)

func CreateRarity(r *models.Rarity) (*app.GwentapiRarity, error) {
	uuid := helpers.EncodeUUID(r.UUID)

	result := &app.GwentapiRarity{
		Name: r.Name,
		Href: helpers.RarityURL(uuid),
		UUID: uuid,
	}

	return result, nil
}

func CreateRarityLink(r *models.Rarity) (*app.GwentapiRarityLink, error) {
	uuid := helpers.EncodeUUID(r.UUID)
	rarityLink := &app.GwentapiRarityLink{
		Href: helpers.RarityURL(uuid),
		Name: r.Name,
	}
	return rarityLink, nil
}
