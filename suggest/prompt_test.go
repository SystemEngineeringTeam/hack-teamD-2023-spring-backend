package suggest

import (
	"log"
	"os"
	"testing"
	"text/template"
)

func TestTmpl(t *testing.T) {
	Tmpl("a")
}
func Tmpl(str string) {
	tmpl, err := template.New("main").Parse(PromptV1)
	if err != nil {
		log.Panic(err)
	}
	tmpl.Execute(os.Stdout, str)
}
