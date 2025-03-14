package responses

import (
	"github.com/baldurstod/go-printful-api-model/schemas"
	printfulmodel "github.com/baldurstod/go-printful-sdk/model"
)

type Order struct {
	Code   int           `json:"code" bson:"code" mapstructure:"code"`
	Result schemas.Order `json:"result" bson:"result" mapstructure:"result"`
}

type CreateOrderResponse struct {
	Success bool `json:"success" bson:"success" mapstructure:"success"`
	Result  struct {
		Order printfulmodel.Order `json:"order" bson:"order" mapstructure:"order"`
	} `json:"result" bson:"result" mapstructure:"result"`
	Error `json:"error" bson:"error" mapstructure:"error"`
}
