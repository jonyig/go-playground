package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handler struct {
	remoteHost *url.URL
}

func GetRemoteUrl() *url.URL {
	remote, err := url.Parse("http://localhost:8081/")
	if err != nil {
		panic(err)
	}

	return remote
}
func CreateRouter(r *gin.Engine, h *handler) {
	r.Any("/*proxyPath", h.ProxyHandler)
}

func CreateHandler(r *gin.Engine, remoteHost *url.URL) *handler {
	h := &handler{
		remoteHost: remoteHost,
	}

	CreateRouter(r, h)
	r.Any("/*proxyPath", h.ProxyHandler)

	return h

}

func (h *handler) ProxyHandler(c *gin.Context) {
	proxy := httputil.NewSingleHostReverseProxy(h.remoteHost)
	proxy.Director = func(req *http.Request) {
		//b, _ := ioutil.ReadAll(req.Body)
		//fmt.Println(string(b))
		req.Host = h.remoteHost.Host
		req.URL.Scheme = h.remoteHost.Scheme
		req.URL.Host = h.remoteHost.Host
		//...£
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
