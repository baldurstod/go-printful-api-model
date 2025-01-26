package responses

import (
	printfulmodel "github.com/baldurstod/go-printful-sdk/model"
)

type GetProductResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Product  printfulmodel.Product   `json:"product"`
		Variants []printfulmodel.Variant `json:"variants"`
	} `json:"result"`
	Error `json:"error"`
}
