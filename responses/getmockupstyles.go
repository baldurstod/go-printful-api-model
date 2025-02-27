package responses

import (
	printfulmodel "github.com/baldurstod/go-printful-sdk/model"
)

type GetMockupStylesResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Styles []printfulmodel.MockupStyles `json:"styles"`
	} `json:"result"`
	Error `json:"error"`
}
