package requests

import "github.com/baldurstod/go-printful-sdk/model"

type CalculateShippingRates struct {
	Recipient model.ShippingRatesAddress                 `mapstructure:"recipient"`
	Items     []model.CatalogOrWarehouseShippingRateItem `mapstructure:"items"`
	Currency  string                                     `mapstructure:"currency"`
	Locale    string                                     `mapstructure:"locale"`
}
