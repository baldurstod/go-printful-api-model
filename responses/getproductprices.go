package responses

import (
	printfulmodel "github.com/baldurstod/go-printful-sdk/model"
)

type GetProductPricesResponse struct {
	Success bool                        `json:"success"`
	Result  printfulmodel.ProductPrices `json:"result"`
	Error   `json:"error"`
}
