package db

import (
	"context"
	"fmt"
	"github.com/vynious/gd-telemessenger-ms/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	dbname   string
	collname string
)

type Repository struct {
	mg       *mongo.Client
	timeout  time.Duration
	dbname   types.DatabaseName
	collname types.CollectionName
}

func SpawnMongoClient(cfg types.MongoConfig) (*Repository, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(string(cfg.Url)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}
	return &Repository{
		mg:       client,
		timeout:  time.Duration(2) * time.Second,
		dbname:   cfg.DBName,
		collname: cfg.CollName,
	}, nil

}

func (r *Repository) CreateSubscription(th types.TelegramHandle, cid types.ChatID) error {
	_, err := r.mg.Database(dbname).Collection(collname).InsertOne(context.TODO(), bson.D{
		{"TelegramHandle", th},
		{"ChatId", cid}},
	)
	if err != nil {
		return fmt.Errorf("error creating subscription: %w", err)
	}
	return nil
}

func (r *Repository) GetSubscription(th types.TelegramHandle) (types.ChatID, error) {
	var user types.UserDocument
	if err := r.mg.Database(dbname).Collection(collname).FindOne(context.TODO(), bson.D{
		{"TelegramHandle", th},
	}).Decode(&user); err != nil {
		return "", fmt.Errorf("error getting subscription: %w", err)
	}
	return user.ChatId, nil
}
