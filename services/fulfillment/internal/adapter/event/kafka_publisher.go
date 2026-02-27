package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rengumin/fulfillment/internal/core/entity"
	"github.com/segmentio/kafka-go"
)

type kafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(brokers []string) *kafkaPublisher {
	return &kafkaPublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *kafkaPublisher) publish(ctx context.Context, topic entity.Topic, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	return p.writer.WriteMessages(ctx, kafka.Message{
		Topic: string(topic),
		Key:   []byte(key),
		Value: data,
	})
}

func (p *kafkaPublisher) PublishPickupScheduled(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	event := map[string]interface{}{
		"event_type":          entity.EventPickupScheduled,
		"package_number":      pkg.PackageNumber,
		"order_id":            pkg.OrderID,
		"shop_id":             pkg.ShopID,
		"pickup_scheduled_at": pkg.PickupScheduledAt,
	}
	return p.publish(ctx, entity.TopicPickupScheduled, pkg.PackageNumber, event)
}

func (p *kafkaPublisher) PublishPickedUp(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	event := map[string]interface{}{
		"event_type":     entity.EventPickedUp,
		"package_number": pkg.PackageNumber,
		"order_id":       pkg.OrderID,
		"shop_id":        pkg.ShopID,
		"picked_up_at":   pkg.PickupCompletedAt,
	}
	return p.publish(ctx, entity.TopicPickedUp, pkg.PackageNumber, event)
}

func (p *kafkaPublisher) PublishInTransit(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	event := map[string]interface{}{
		"event_type":       entity.EventInTransit,
		"package_number":   pkg.PackageNumber,
		"order_id":         pkg.OrderID,
		"current_location": pkg.CurrentHubLocation,
		"updated_at":       pkg.UpdatedAt,
	}
	return p.publish(ctx, entity.TopicInTransit, pkg.PackageNumber, event)
}

func (p *kafkaPublisher) PublishOutForDelivery(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	event := map[string]interface{}{
		"event_type":         entity.EventOutForDelivery,
		"package_number":     pkg.PackageNumber,
		"order_id":           pkg.OrderID,
		"estimated_delivery": pkg.EstimatedDelivery,
	}
	return p.publish(ctx, entity.TopicOutForDelivery, pkg.PackageNumber, event)
}

func (p *kafkaPublisher) PublishDelivered(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	event := map[string]interface{}{
		"event_type":     entity.EventDelivered,
		"package_number": pkg.PackageNumber,
		"order_id":       pkg.OrderID,
		"delivered_at":   pkg.DeliveredAt,
	}
	return p.publish(ctx, entity.TopicDelivered, pkg.PackageNumber, event)
}

func (p *kafkaPublisher) PublishDeliveryFailed(ctx context.Context, pkg *entity.FulfillmentPackage) error {
	event := map[string]interface{}{
		"event_type":           entity.EventDeliveryFailed,
		"package_number":       pkg.PackageNumber,
		"order_id":             pkg.OrderID,
		"failure_reason":       pkg.DeliveryFailureReason,
		"delivery_attempts":    pkg.DeliveryAttempts,
		"max_attempts_reached": pkg.DeliveryAttempts >= 3,
	}
	return p.publish(ctx, entity.TopicDeliveryFailed, pkg.PackageNumber, event)
}

func (p *kafkaPublisher) Close() error {
	return p.writer.Close()
}

