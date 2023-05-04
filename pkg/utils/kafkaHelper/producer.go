package kafkaHelper

import "github.com/Shopify/sarama"

type kafkaProducer struct {
	config     *sarama.Config
	brokerList []string
}

func NewKafkaProducer(brokerList []string, config *sarama.Config) kafkaProducer {
	return kafkaProducer{
		config:     config,
		brokerList: brokerList,
	}
}

func (k kafkaProducer) GetSyncProducer() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(k.brokerList, k.config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func (k kafkaProducer) SyncProducerSendMessage(topic, key, message string) (partition int32, offset int64, err error) {
	producer, err := k.GetSyncProducer()
	if err != nil {
		return 0, 0, err
	}
	defer func() {
		if err = producer.Close(); err != nil {
			return
		}
	}()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err = producer.SendMessage(msg)
	if err != nil {
		return 0, 0, err
	}
	return
}

func (k kafkaProducer) NewAsyncProducer() (sarama.AsyncProducer, error) {
	producer, err := sarama.NewAsyncProducer(k.brokerList, k.config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func (k kafkaProducer) AsyncProducerSendMessage(topic, key, message string) (err error) {
	producer, err := k.NewAsyncProducer()
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}

	producer.Input() <- msg
	defer producer.AsyncClose()

	return
}

func (k kafkaProducer) AsyncProducerSendBatchMessage(topic, key string, message []string) (err error) {
	producer, err := k.NewAsyncProducer()
	if err != nil {
		return err
	}
	defer func() {
		if err = producer.Close(); err != nil {
			return
		}
	}()
	for _, v := range message {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(v),
		}

		producer.Input() <- msg
	}

	return
}
