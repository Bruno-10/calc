package calc

// Result represents the Execute API web response.
type Result struct {
	Total    float64   `json:"total"`
	SumGroup []float64 `json:"sumGroup"`
}
