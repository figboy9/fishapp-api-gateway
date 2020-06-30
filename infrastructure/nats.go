package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func NewNatsStreamingConn() (stan.Conn, error) {
	clientID := uuid.New().String()
	log.Printf("nats clientID is %s", clientID)

	return stan.Connect(conf.C.Nats.ClusterID, "fishapp-api-gateway-"+clientID, stan.NatsURL(conf.C.Nats.URL))
}
