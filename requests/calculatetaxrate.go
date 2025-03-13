package requests

import (
	"github.com/baldurstod/go-printful-api-model/schemas"
)

type CalculateTaxRate struct {
	Recipient schemas.TaxAddressInfo `json:"recipient" bson:"recipient" mapstructure:"recipient"`
}
