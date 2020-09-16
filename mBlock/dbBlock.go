package mBlock

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDBBlock struct {
	Block
}

func (mongodbBlock *MongoDBBlock) Connect() *mongo.Client {
	uri := "mongodb+srv://127.0.0.1"
	ctx ,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err !=nil{
		fmt.Println(err)
	}
	return client
}

