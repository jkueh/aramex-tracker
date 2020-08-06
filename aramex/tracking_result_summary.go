package aramex

// TrackingResultSummary is the data that we get back from the /tracking-api endpoint.
type TrackingResultSummary struct {
	Result      TrackingResult `json:"result"`
	GeneratedIn string         `json:"generated_in"`

	// These fields are present in an error payload (e.g. if you provide a bad tracking number)
	Error string `json:"error"`
	// Not too sure what this is
	Data interface{} `json:"data"`
}
