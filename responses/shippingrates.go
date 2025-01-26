package responses

import (
	"github.com/baldurstod/go-printful-api-model/schemas"
)

type ShippingRates struct {
	Code   int                    `json:"code" bson:"code" mapstructure:"code"`
	Result []schemas.ShippingInfo `json:"result" bson:"result" mapstructure:"result"`
}
