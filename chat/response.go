package chat

// Response : Store Response from OpenAI API
type Response struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Usage   *Usage    `json:"usage"`
	Choices []*Choice `json:"choices"`
}

// Usage : Store Usage of OpenAI API
type Usage struct {
	PromptTokens int `json:"prompt_tokens"`

	CompletionTokens int `json:"completion_tokens"`

	TotalTokens int `json:"total_tokens"`
}

// Choice : Store Choice of OpenAI API
type Choice struct {
	Message *ResponseMessage `json:"message"`

	FinishReason string `json:"finish_reason"`

	Index int `json:"index"`
}

type ResponseMessage struct {
	Role string `json:"role"`

	Content string `json:"content"`
}
