//go:build debug
// +build debug

package web

import "github.com/gin-gonic/gin"

func setMode() {
	gin.SetMode(gin.DebugMode)
}
