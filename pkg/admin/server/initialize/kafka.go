package initialize

import (
	"context"
	"fmt"
	"log"
	"thanhbk113/internal/config"

	"github.com/logrusorgru/aurora"
	"github.com/segmentio/kafka-go"
)

var (
	conn *kafka.Conn
)

func initKafka() {
	broker := config.GetConfig().KAFKA_URI

	connKafka, err := kafka.Dial("tcp", broker)

	if err != nil {
		fmt.Println(aurora.Red("*** Kafka connection failed:" + broker + " ***"))
		panic(err)
	}

	conn = connKafka

	fmt.Println(aurora.Green("*** Kafka connected successfully:" + broker + " ***"))
}

// GetKafkaConn return kafka connection
func GetKafkaConn() *kafka.Conn {
	return conn
}

// Create topic kafka
func CreateTopic(topic string) {
	err := conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})

	if err != nil {
		fmt.Println(aurora.Red("*** Create topic failed:" + topic + " ***"))
		panic(err)
	}

	fmt.Println(aurora.Green("*** Create topic successfully:" + topic + " ***"))
}

func DeleteTopic(topic string) {
	err := conn.DeleteTopics(topic)

	if err != nil {
		fmt.Println(aurora.Red("*** Delete topic failed:" + topic + " ***"))
		panic(err)
	}

	fmt.Println(aurora.Green("*** Delete topic successfully:" + topic + " ***"))
}

// send message to topic kafka
func SendMessage(topic string, message string, key string) {

	w := kafka.Writer{
		Addr:     kafka.TCP(config.GetConfig().KAFKA_URI),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),

		kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	fmt.Println(aurora.Green("*** Send message:" + message + "with key:" + key + "successfully to topic:" + topic + " ***"))
}

// MatchMessage match message with key
func MatchMessage(ctx context.Context, topic string, key string) bool {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.GetConfig().KAFKA_URI},
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		if string(m.Key) == key {
			fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			return true
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

	return false
}

// ListTopic list all topic
func ListTopic() {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	i := 0
	for k := range m {
		fmt.Println("topic kafka:", i, "->", k)
		i++
	}

}
