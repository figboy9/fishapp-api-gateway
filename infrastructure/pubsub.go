package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/ezio1119/fishapp-api-gateway/conf"
)

func NewCloudPubSubClient() *pubsub.Client {
	c, err := pubsub.NewClient(context.Background(), conf.C.Gcp.ProjectID)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
