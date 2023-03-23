package chat

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

type RequestData struct {
	Model    string            `json:"model"`
	Messages []*RequestMessage ` json:"messages"`
	//MaxToken int               `json:"max_token"`
	MaxToken int `json:"-"`
}

func NewRequest(modelName string, messages []*RequestMessage, maxToken int) *RequestData {
	return &RequestData{
		Model:    modelName,
		Messages: messages,
		MaxToken: maxToken,
	}
}

type RequestMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

func NewRequestMessage(role string, content string) *RequestMessage {
	return &RequestMessage{
		Content: content,
		Role:    role,
	}
}
