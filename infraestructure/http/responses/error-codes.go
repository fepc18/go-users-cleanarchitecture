package responses

type ErrorCode string

const (
	ExpectedField      ErrorCode = "ExpectedField"
	MissingField       ErrorCode = "MissingField"
	InvalidField       ErrorCode = "InvalidField"
	InvalidCredentials ErrorCode = "InvalidCredentials"
	Unexpected         ErrorCode = "Unexpected"
)

/*
func (e ErrorCode) String() string {
	return [...]string{"ExpectedField", "MissingField", "InvalidField", "InvalidCredentials", "Unexpected"}[e-1]
}

func (e ErrorCode) EnumIndex() int {
	return int(e)
}*/
