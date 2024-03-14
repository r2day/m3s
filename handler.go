package m3s

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const (
	defaultDB = "m3s"
	dbNameKey = "M3S_DB"
)

var (
	// MDB 设置全局数据库handler
	MDB *mongo.Database
)

// New 创建新连接池
func New(uri string) {
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv(dbNameKey)
	if dbName == "" {
		dbName = defaultDB
	}

	MDB = client.Database(dbName)
}
