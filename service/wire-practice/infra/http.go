package infra

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}
