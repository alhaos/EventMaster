package apiServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Config struct {
	port int
}

// New ...
func New() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("html/templates/*")

	//authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
	//	"user1": "love",
	//	"user2": "god",
	//	"user3": "sex",
	//}))

	router.POST("/", HeaderRoot)
	return router

}

// HeaderRoot ...
func HeaderRoot(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		c.PostForm("a"),
	)

}
