package app

import (
	"github.com/daniel-888/BloXroute-CS/models"
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	conn *amqp.Connection
	config Configurations
}

func New(config Configurations) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) InitClient() error {
	var err error
	c.conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s", c.config.RabbitMQConfig.User, c.config.RabbitMQConfig.Password, c.config.RabbitMQConfig.URL))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Cleanup() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) SendCommand(ctx context.Context, command *models.Command) error {
	log.WithField(traceIDkey, ctx.Value(traceIDkey)).
		Info("Sending command. Type: ", command.Type.String(), ", Payload: ", command.String())
	
	ch, err := c.conn.Channel()

	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(
		c.config.RabbitMQConfig.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body, err := proto.Marshal(command)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			Headers: map[string]any{
				traceIDkey: ctx.Value(traceIDkey),
			},
			DeliveryMode: amqp.Persistent,
			ContentType: 	"application/protobuf",
			Body: 				body,
		},
	)
	if err != nil {
		return err
	}

	log.WithField(traceIDkey, ctx.Value(traceIDkey)).Info("[x] Sent message")
	return nil
}