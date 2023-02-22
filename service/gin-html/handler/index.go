package handler

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Main website",
		"ttt":   "Main website12313",
	})

}

func CreateMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "service/gin-html/template/base.html", "service/gin-html/template/index.html")
	//r.AddFromFiles("archives", "service/gin-html/template/base.html", "service/gin-html/template/archives.html")

	return r
}
