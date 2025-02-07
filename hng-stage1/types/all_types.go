package types

// Response structure for number classification
type Response struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

// ErrorResponse structure for invalid input
type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}
