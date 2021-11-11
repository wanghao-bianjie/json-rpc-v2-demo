package types

type (
	EchoParams struct {
		Name string `json:"name"`
	}
	EchoResult struct {
		Message string `json:"message"`
	}

	PositionalParams []int
	PositionalResult struct {
		Message []int `json:"message"`
	}
)
