package web

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func Router(server Server) *gin.Engine {

	router := gin.Default()
	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		}))
	*/
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	router.GET("/guest", GuestAccess)
	router.NoRoute(NotImplemented)
	return router
}

//router.GET("/nologin", func(c *gin.Context) {
//c.JSONP(200, g)
//chatEntity := suggest.GenOneWord(c.Request.Form.Get("q"), suggest.PromptV1)
//chatEntity := suggest.GenOneWord(c.Query("q"), suggest.PromptV1)
//var wg *sync.WaitGroup
//resultChann, errChann := chatEntity.Request(wg)
//var list []string
//var err error
//log.Print("tet")
//func() {
//	for {
//		select {
//		case err, errOpen := <-errChann:
//			if errOpen {
//				log.Print(err)
//				return
//			}
//		case v, ok := <-resultChann:
//			if ok {
//				list = append(list, v)
//			} else {
//				return
//			}
//
//		}
//	}
//}()
//if err != nil {
//	log.Print(err)
//	c.AbortWithStatus(http.StatusInternalServerError)
//
//}
//resultJSON, err := json.Marshal(list)
////c.String(http.StatusOK,resultJSON.)
//if _, err = io.Copy(c.Writer, bytes.NewReader(resultJSON)); err != nil {
//	c.AbortWithStatus(http.StatusInternalServerError)
//}
//response, err := chatEntity.Request()
//if err != nil {
//	c.AbortWithStatus(http.StatusInternalServerError)
//	return
//}
//c.JSON(http.StatusOK, response)
//})
