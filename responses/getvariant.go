package responses

import (
	printfulmodel "github.com/baldurstod/go-printful-sdk/model"
)

type GetVariantResponse struct {
	Success bool                  `json:"success"`
	Result  printfulmodel.Variant `json:"result"`
	Error   `json:"error"`
}
