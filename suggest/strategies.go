package suggest

import (
	"log"
	"os"
	"time"

	"github.com/SystemEngineeringTeam/hack-teamD-2023-spring-backend/chat"
)

type Request string

type Strategy interface {
	Request() (string, error)
	debug()
	GetRaw() string
}
type oneWord struct {
	src          any
	templ        string
	ResponseText string
}

var preset = chat.NewChatCompletions("gpt-3.5-turbo", os.Getenv("GPT_KEY"),
	1500, 10*time.Minute)

func (w oneWord) Request() (string, error) {
	// errChan := make(chan error, 1)
	// wg.Add(1)
	cmp := preset

	res, err := cmp.AskOneQuestion(GenQuery(w.templ, w.src))
	if err != nil {
		return "", err
	}
	if res == nil {
		panic(err)
	}
	var suggest string
	if len(res.Choices) == 0 {
		panic("Choice is nil")
		log.Panic("Choice is nil")

	}
	log.Printf("%+v \n", res.Choices)
	for _, v := range res.Choices {
		log.Printf("v: %+v \n", v)
		log.Println("v.M: ", v.Message)
		w.ResponseText = MessageExtract(v.Message)
		log.Println("MessageExtract: ", w.ResponseText)
		if w.ResponseText == "" {
			log.Panic("void string")
		}
		log.Println("Message: ", w.ResponseText)
		suggest = suggest + ExtractSliceFromStr(w.ResponseText)
		// responses=append(responses, v.Message.Content)
	}
	return suggest, nil
	//responses := make(chan string, 1)
	/*
		go func() {
			//defer wg.Done()
			defer close(responses)
			defer close(errChan)
			r := bytes.Buffer{}
			var t template.Template
			if t, err := template.New("templ").Parse(w.templ); t != nil {
				errChan <- err
				return
			}
			if err := t.Execute(&r, w.originWord); err != nil {
				errChan <- err
				return
			}
			var resp *chat.Response
			var err error
			if resp, err = cmp.AskOneQuestion(r.String()); err != nil {
				errChan <- err
				return
			}

			for _, v := range resp.Choices {
				//rows := strings.Split(v.Message.Content, "\n")
				//last := rows[len(rows)-1]
				//for _, v2 := range strings.Split(last, ",") {
				//	responses <- v2
				//}
				index := strings.Index(
					v.Message.Content, "[")
				lastIndex := strings.LastIndex(v.Message.Content, "]")
				if index > 0 && lastIndex > index {
					responses <- v.Message.Content[index:lastIndex]
				}
			}
			//strings.Index()
		}()

		//return responses, errChan
		return*/
}

func GenOneWord(s any, templ string) Strategy {
	if s == "" {
		return nil
	}
	//if err != nil {
	//	return nil
	//}
	return oneWord{
		src:   s,
		templ: templ,
	}
}

func (w oneWord) debug() {
	println(w.src)
	println(w.templ)
}

func (w oneWord) GetRaw() string {
	return w.ResponseText
}
