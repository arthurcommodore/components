package mongo

import (
	"context"

	drivermongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client   *drivermongo.Client
	database *drivermongo.Database
}

func Connect(ctx context.Context, uri string, databaseName string) (*Client, error) {
	client, err := drivermongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil, err
	}

	return &Client{
		client:   client,
		database: client.Database(databaseName),
	}, nil
}

func (c *Client) Database() *drivermongo.Database {
	return c.database
}

func (c *Client) Close(ctx context.Context) error {
	if c == nil || c.client == nil {
		return nil
	}

	return c.client.Disconnect(ctx)
}
