package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyHandler(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8081/")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		//b, _ := ioutil.ReadAll(req.Body)
		//fmt.Println(string(b))
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		//...£
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
