package entity

// Topic defines Kafka topic names for fulfillment events.
type Topic string

const (
	TopicPickupScheduled Topic = "fulfillment.pickup_scheduled"
	TopicPickedUp        Topic = "fulfillment.picked_up"
	TopicInTransit       Topic = "fulfillment.in_transit"
	TopicOutForDelivery  Topic = "fulfillment.out_for_delivery"
	TopicDelivered       Topic = "fulfillment.delivered"
	TopicDeliveryFailed  Topic = "fulfillment.delivery_failed"
)

// EventType defines the event_type field values inside Kafka messages.
type EventType string

const (
	EventPickupScheduled EventType = "fulfillment.pickup_scheduled"
	EventPickedUp        EventType = "fulfillment.picked_up"
	EventInTransit       EventType = "fulfillment.in_transit"
	EventOutForDelivery  EventType = "fulfillment.out_for_delivery"
	EventDelivered       EventType = "fulfillment.delivered"
	EventDeliveryFailed  EventType = "fulfillment.delivery_failed"
)
