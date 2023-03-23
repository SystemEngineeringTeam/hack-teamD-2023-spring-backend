package chat

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestChatCompletions_MockSendMessage(t *testing.T) {
	r := SimpleHTTPServer()
	//Caution: this test may use API Usage
	content := "りんごに関係する単語を配列として出力\n"
	c := NewChatCompletions("gpt-3.5-turbo", os.Getenv("GPT_KEY"),
		1500, 60*time.Second)
	messages := []*RequestMessage{
		NewRequestMessage("user", content),
	}
	go r.Run(":8081")
	res, err := c.MockSendMessage(messages)
	println(res)

	println(err)
}

func TestChatCompletions_SendMessage(t *testing.T) {
	//r := SimpleHTTPServer()
	//Caution: this test may use API Usage
	content := "りんごに関係する単語を配列として出力\n"
	c := NewChatCompletions("gpt-3.5-turbo", os.Getenv("GPT_KEY"),
		1500, 60*time.Second)
	messages := []*RequestMessage{
		NewRequestMessage("user", content),
	}
	res, err := c.SendMessage(messages)
	println(res.Choices)
	for _, v := range res.Choices {
		println(v.Message.Content)

	}

	println(err)
}
func (c ChatCompletions) MockSendMessage(messages []*RequestMessage) (*Response, error) {
	data, err := json.Marshal(NewRequest(c.model, messages, c.maxTokens))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:8081/v1/chat/completions",
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.secret)

	client := &http.Client{
		// リソース節約のためにタイムアウトを設定する
		Timeout: 20 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("bad status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func SimpleHTTPServer() *gin.Engine {
	r := gin.Default()
	r.NoRoute(func(context *gin.Context) {
		println("===Header===")
		println("Authorization: ", context.Request.Header.Get("Authorization"))
		println("Content-Type: ", context.ContentType())
		println("===Body===")
		io.Copy(os.Stdout, context.Request.Body)
		context.String(http.StatusOK, "{{.}}",
			//{"id":"chatcmpl-6u8z10YYeGZmJq7loeTxjS8muVPD4","object":"chat.completion","created":1678838919,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":26,"completion_tokens":106,"total_tokens":132},"choices":[{"message":{"role":"assistant","content":"\n\n[\"りんご\", \"果物\", \"林檎\", \"落花生\", \"リンゴ酸\", \"パイ\", \"フルーツ\", \"ジュース\", \"サイダー\", \"キャラメル\", \"保存料\", \"甘酸っぱい\", \"紅玉\", \"青森\", \"山形\", \"長野\", \"中津川\"]"},"finish_reason":"stop","index":0}]}
			"{\"id\":\"chatcmpl-6u8z10YYeGZmJq7loeTxjS8muVPD4\",\"object\":\"chat.completion\",\"created\":1678838919,\"model\":\"gpt-3.5-turbo-0301\",\"usage\":{\"prompt_tokens\":26,\"completion_tokens\":106,\"total_tokens\":132},\"choices\":[{\"message\":{\"role\":\"assistant\",\"content\":\"\\n\\n[\\\"りんご\\\",\\\"果物\\\",\\\"林檎\\\",\\\"落花生\\\",\\\"リンゴ酸\\\",\\\"パイ\\\",\\\"フルーツ\\\",\\\"ジュース\\\",\\\"サイダー\\\",\\\"キャラメル\\\",\\\"保存料\\\",\\\"甘酸っぱい\\\",\\\"紅玉\\\",\\\"青森\\\",\\\"山形\\\",\\\"長野\\\",\\\"中津川\\\"]\"},\"finish_reason\":\"stop\",\"index\":0}]}",
		)
	})
	return r
}
