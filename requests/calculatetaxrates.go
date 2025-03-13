package requestbodies

import (
	"github.com/baldurstod/go-printful-api-model/schemas"
)

type CalculateTaxRates struct {
	Recipient schemas.TaxAddressInfo `json:"recipient" bson:"recipient" mapstructure:"recipient"`
}
