package service

import (
	"context"
	"workflow-service/service/errors"
	"workflow-service/service/util"

	dapr "github.com/dapr/go-sdk/client"
)

// A wrapper service for a Dapr client.
type DaprService struct {
	DaprClient dapr.Client
}

// A constructor function for DaprService structure.
func NewDapr(client dapr.Client) *DaprService {
	return &DaprService{
		DaprClient: client,
	}
}

// Method for publishing an event in the system.
func (srvc DaprService) PublishEvent(ctx context.Context, topic string, data interface{}) error {
	if srvc.DaprClient == nil {
		return errors.ErrDaprIsNotInitialized
	}
	return srvc.DaprClient.PublishEvent(ctx, util.PUBSUB_NAME, topic, data)
}
