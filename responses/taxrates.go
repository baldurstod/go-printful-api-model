package responses

import (
	"github.com/baldurstod/go-printful-api-model/schemas"
)

type TaxRates struct {
	Code   int             `json:"code" bson:"code" mapstructure:"code"`
	Result schemas.TaxInfo `json:"result" bson:"result" mapstructure:"result"`
}
