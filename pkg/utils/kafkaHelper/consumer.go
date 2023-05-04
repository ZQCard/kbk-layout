package kafkaHelper

import (
	"context"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	consumer    sarama.ConsumerGroup
	dataChannel chan string
	topic       string
	groupID     string
}

func NewKafkaConsumer(config *sarama.Config, brokerList []string, topic string, groupID string) (*KafkaConsumer, error) {
	consumer := KafkaConsumer{
		dataChannel: make(chan string),
		topic:       topic,
		groupID:     groupID,
	}

	var err error
	consumer.consumer, err = sarama.NewConsumerGroup(brokerList, groupID, config)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (k *KafkaConsumer) Consume() error {
	handler := consumerHandler{k.dataChannel}
	topics := []string{k.topic}

	for {
		err := k.consumer.Consume(context.Background(), topics, &handler)
		if err != nil {
			return err
		}
	}
}

func (k *KafkaConsumer) GetDataChannel() chan string {
	return k.dataChannel
}

func (k *KafkaConsumer) Close() {
	if err := k.consumer.Close(); err != nil {
		return
	}
}

type consumerHandler struct {
	dataChannel chan string
}

func (h *consumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		data := string(message.Value)
		h.dataChannel <- data
		session.MarkMessage(message, "")
	}
	return nil
}
