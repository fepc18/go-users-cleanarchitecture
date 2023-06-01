package responses

type Error struct {
	ErrorCode ErrorCode `json:"errorCode"`
	Message   string    `json:"message"`
	//	Path    string `json:"path"`
	//Utl     interface{} `json:"url"`
}

type ErrorResponse struct {
	Code    int     `json:"-"`
	Id      string  `json:"id"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors,omitempty"`
}

func NewErrorResponse(id string, code int, message string, errors []Error) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Id:      id,
		Errors:  errors,
	}
}

// add errors
func (err *ErrorResponse) AddError(errorapi Error) {
	err.Errors = append(err.Errors, errorapi)
}
