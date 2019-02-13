package data

type PipelineRequest struct {
	Id        string
	Source    string
	Input     string
	Entities  []string
	Sentiment float64
	Error     error
}
