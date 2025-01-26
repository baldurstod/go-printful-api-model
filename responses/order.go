package responses

import (
	"github.com/baldurstod/go-printful-api-model/schemas"
)

type Order struct {
	Code   int           `json:"code" bson:"code" mapstructure:"code"`
	Result schemas.Order `json:"result" bson:"result" mapstructure:"result"`
}
