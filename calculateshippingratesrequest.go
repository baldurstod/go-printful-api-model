package model

import (
	printfulmodel "github.com/baldurstod/go-printful-sdk/model"
)

type CalculateShippingRatesRequest struct {
	Recipient printfulmodel.ShippingRatesAddress                 `json:"recipient" bson:"recipient" mapstructure:"recipient"`
	Items     []printfulmodel.CatalogOrWarehouseShippingRateItem `json:"items" bson:"items" mapstructure:"items"`
	Currency  string                                             `json:"currency" bson:"currency" mapstructure:"currency"`
	Locale    string                                             `json:"locale" bson:"locale" mapstructure:"locale"`
}
