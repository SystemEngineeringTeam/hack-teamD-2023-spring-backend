package suggest

import (
	"bytes"
	"log"
	"text/template"
)

const PromptV1 = "{{range .}}「{{- . -}}」{{end}}" +
	"に関係する単語3つを配列として出力。配列は以下の形式です\n[\"A\",\"B\"]"

type ParamStruct struct {
	Queries []string `form:"q[]" url:"q[]"`
}

func GenQuery(templ string, src any) string {
	var t *template.Template
	var err error
	if t, err = template.New("templ").Parse(templ); err != nil {
		log.Panic(err)
	}
	r := bytes.Buffer{}
	if err := t.ExecuteTemplate(&r, "templ", src); err != nil {
		log.Panic(err)
	}
	return r.String()
}
