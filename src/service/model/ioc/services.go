package ioc

import "context"

// An interface for a Dapr service.
type IDaprService interface {
	// Method for publishing an event in the system.
	PublishEvent(ctx context.Context, topic string, data interface{}) error
}
