package suggest

import (
	"strings"

	"github.com/SystemEngineeringTeam/hack-teamD-2023-spring-backend/chat"
)

func MessageExtract(message *chat.ResponseMessage) string {
	if message == nil {
		panic("nil pointer")
	}
	return message.Content
}

func ExtractSliceFromStr(s string) string {
	index := strings.Index(
		s, "[")
	lastIndex := strings.LastIndex(s, "]")
	if index > -1 && lastIndex >= index {
		return s[index : lastIndex+1]
	}
	return ""
}
