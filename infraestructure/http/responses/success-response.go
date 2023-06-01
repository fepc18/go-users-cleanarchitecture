package responses

type SuccessResponse struct {
	Data  interface{} `json:"data"`
	Links interface{} `json:"links"`
}

func NewSuccessResponse(data interface{}, links interface{}) *SuccessResponse {
	return &SuccessResponse{
		Data:  data,
		Links: links,
	}
}
