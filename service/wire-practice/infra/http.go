package infra

import "github.com/google/wire"

type HttpClient struct{}

func GetHttpClient() *HttpClient {
	return &HttpClient{}
}

var HttpProviderSet = wire.NewSet(GetHttpClient)
