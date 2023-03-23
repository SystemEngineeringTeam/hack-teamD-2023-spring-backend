package chat

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type ChatCompletions struct {
	timeout time.Duration

	// 応答の最大トークン
	maxTokens int

	// チャットに使用するモデルのID
	model string

	// APIキー
	secret string
}
type Completion interface {
	SendMessage([]*RequestMessage) (*Response, error)
}

func NewChatCompletions(
	model string,
	secret string,
	maxTokens int,
	timeout time.Duration,
) *ChatCompletions {
	//) *Completion {
	return &ChatCompletions{
		maxTokens: maxTokens,
		model:     model,
		secret:    secret,
		timeout:   timeout,
	}
}

// APIにメッセージを送信する
func (c ChatCompletions) SendMessage(messages []*RequestMessage) (*Response, error) {
	data, err := json.Marshal(NewRequest(c.model, messages, c.maxTokens))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/chat/completions",
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.secret)

	client := &http.Client{
		// リソース節約のためにタイムアウトを設定する
		Timeout: 20 * time.Minute,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("bad status: " + resp.Status)
		return nil, errors.New("bad status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println("response: " + bytes.NewBuffer(body).String())

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// 一回きりの質問をする
func (c ChatCompletions) AskOneQuestion(content string) (*Response, error) {
	messages := []*RequestMessage{
		NewRequestMessage("user", content),
	}
	return c.SendMessage(messages)
}
