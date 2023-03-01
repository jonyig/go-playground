package infra

type MongoClient struct{}

func NewMongoClient() *MongoClient {
	return &MongoClient{}
}
