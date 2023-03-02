package infra

import "github.com/google/wire"

type MongoClient struct{}

func GetMongoClient() *MongoClient {
	return &MongoClient{}
}

var MongoProviderSet = wire.NewSet(GetMongoClient)
