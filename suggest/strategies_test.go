package suggest

import (
	"log"
	"strings"
	"testing"

	"github.com/SystemEngineeringTeam/hack-teamD-2023-spring-backend/chat"
)

func TestGenOneWord(t *testing.T) {
	a := GenOneWord("apple", PromptV1)
	a.debug()
}

func TestOneWord_Request(t *testing.T) {
	// a := GenOneWord("apple", PromptV1)
	// response, err := a.Request()

	//if err != nil {
	//	log.Panic(err)
	//}
	response := chat.Response{
		ID:      "chatcmpl",
		Object:  "chat.completion",
		Created: 1678838984,
		Model:   "gpt-3.5-turbo-0301",
		Usage: &chat.Usage{
			PromptTokens:     0,
			CompletionTokens: 0,
			TotalTokens:      0,
		},
		Choices: []*chat.Choice{{
			Message: &chat.ResponseMessage{
				Role:    "assistant",
				Content: "[\\\"チャット\\\", \\\"オンライン\\\", \\\"コミュニケーション\\\", \\\"メッセージ\\\", \\\"会話\\\", \\\"テキスト\\\", \\\"ユーザー\\\", \\\"プロファイル\\\", \\\"フレンドリスト\\\", \\\"相手の入力待ち\\\", \\\"エモジ\\\", \\\"スマイリー\\\", \\\"チャットルーム\\\", \\\"グループチャット\\\", \\\"プライベートチャット\\\", \\\"チャットボット\\\"]",
			},
			FinishReason: "stop",
			Index:        0,
		}},
	}
	//println(response)
	//for _, v := range response.Choices {
	//	println(v.Message.Content)
	//}
	var responses []string
	for _, v := range response.Choices {
		index := strings.Index(
			v.Message.Content, "[")
		lastIndex := strings.LastIndex(v.Message.Content, "]")
		println(index)
		println(lastIndex)
		if index > -1 && lastIndex > index {
			responses = append(responses, v.Message.Content[index:lastIndex])
		}
		// responses=append(responses, v.Message.Content)
	}
	println(len(responses))
	println(responses[0])
	/*
		wg := &sync.WaitGroup{}
		ResChan := make(chan string)
		errChan := make(chan error)
		a.Request()

		go func() {
			for {
				select {
				case res, ok := <-ResChan:
					if ok {
						println(res)
					} else {
						break
					}
				case err, ok := <-errChan:
					if ok {
						println(err)
						log.Panic(err)
					} else {
						continue
					}
				}
			}
		}()
		//ResChann, errChann := a.Request(wg)
		println(ResChan)
		wg.Wait()
		println(errChan)
		*
	*/
}

/*
	func (w oneWord) RequestToMock() {
		cmp := preset
		var t *template.Template
		var err error
		if t, err = template.New("templ").Parse(w.templ); t != nil {
			return
		}
	}
*/
func TestOneWord_GenQuery(t *testing.T) {
	a := GenOneWord("apple", PromptV1)
	log.Println(a.GenQuery())
}
