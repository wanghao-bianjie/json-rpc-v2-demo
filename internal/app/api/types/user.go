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

	DetailParams struct {
		Name string `json:"name"`
	}
	DetailResult struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	AddParams struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	AddResult struct {
		Id uint `json:"id"`
	}
)
