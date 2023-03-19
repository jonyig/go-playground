package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_handler_ProxyHandler(t *testing.T) {
	const expected = "this test"
	const expected2 = "this test2"
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/test" {
			w.Write([]byte(expected))
			return
		}

		if r.URL.Path == "/test2" {
			w.Write([]byte(expected2))
			return
		}
		w.Write([]byte(expected2))
	}))
	defer backend.Close()

	backendURL, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatal(err)
	}

	h := handler{remoteHost: backendURL}

	w := CreateTestResponseRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/test?callback=x", nil)

	h.ProxyHandler(c)

	w2 := CreateTestResponseRecorder()
	c2, _ := gin.CreateTestContext(w2)
	c2.Request, _ = http.NewRequest("GET", "/test2?callback=x", nil)
	h.ProxyHandler(c2)

	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, expected2, w2.Body.String())
}

type TestResponseRecorder struct {
	*httptest.ResponseRecorder
	closeChannel chan bool
}

func (r *TestResponseRecorder) CloseNotify() <-chan bool {
	return r.closeChannel
}

func (r *TestResponseRecorder) closeClient() {
	r.closeChannel <- true
}

func CreateTestResponseRecorder() *TestResponseRecorder {
	return &TestResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}
