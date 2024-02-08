package db

import (
	"fmt"
	"github.com/vynious/gd-telemessenger-ms/types"
	"os"
)

func LoadMongoConfig() (types.MongoConfig, error) {
	uri := os.Getenv("MONGO_CONN_URI")
	dbname = os.Getenv("MONGO_DB_NAME")
	collname = os.Getenv("MONGO_COLL_NAME")
	if uri == "" || dbname == "" || collname == "" {
		return types.MongoConfig{}, fmt.Errorf("please check environment variables")
	}

	return types.MongoConfig{
		Url:      types.MongoURI(uri),
		DBName:   types.DatabaseName(dbname),
		CollName: types.CollectionName(collname),
	}, nil
}
