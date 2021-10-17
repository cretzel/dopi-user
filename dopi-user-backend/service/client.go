package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Client struct {
	client *mongo.Client
}

func NewClient() (*Client, error) {
	mongoHost := os.Getenv("MONGODB_HOST")
	mongoPort := os.Getenv("MONGODB_PORT")
	connectionString := "mongodb://" + mongoHost + ":" + mongoPort
	log.Println("Mongo URL:", connectionString)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{client}, err
}

func (c *Client) GetCollection(db string, col string) *mongo.Collection {
	return c.client.Database(db).Collection(col)
}

func (c *Client) Close() {
	if c.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c.client.Disconnect(ctx)
	}
}

func (c *Client) DropDatabase(db string) error {
	if c.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return c.client.Database(db).Drop(ctx)
	}
	return nil
}
