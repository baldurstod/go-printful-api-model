package responses

import (
	printfulmodel "github.com/baldurstod/go-printful-sdk/model"
)

type GetMockupTemplatesResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Templates []printfulmodel.MockupTemplates `json:"templates"`
	} `json:"result"`
	Error `json:"error"`
}
