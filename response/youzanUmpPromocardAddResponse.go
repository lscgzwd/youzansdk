package response

import "github.com/lscgzwd/youzansdk/domain"

// YouzanUmpPromocardAddResponse ..
type YouzanUmpPromocardAddResponse struct {
	baseResponse
	Data struct {
		Promocard domain.Promocard `json:"promocard"`
	} `json:"data"`
}
