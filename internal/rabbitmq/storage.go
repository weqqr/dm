package rabbitmq

import (
	"errors"
	"log"
)

type ServerStorage struct{}

type (
	StreamRequest struct {
		Id string
	}

	StreamResponse struct {
		Result string
	}
)

func (s *ServerStorage) FetchResponse(request StreamRequest) (StreamResponse, error) {
	if request.Id == "" {
		return StreamResponse{}, errors.New("id is empty")
	}

	log.Printf("Fetch response for id : %v", request.Id)

	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(count int64) {
	// 		defer wg.Done()
	// 		time.Sleep(time.Duration(count) * time.Second)
	// 		log.Printf("finishing request number : %v", count)
	// 	}(1)
	// }
	return StreamResponse{
		Result: "success",
	}, nil
}

// func (s *ServerStorage) Receive(config Config) error {
// 	conn, err := amqp.Dial(config.Address)
// 	if err != nil {
// 		return err
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		return err
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"jobs", // name
// 		false,  // durable
// 		false,  // delete when unused
// 		false,  // exclusive
// 		false,  // no-wait
// 		nil,    // arguments
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	var forever chan struct{}

// 	go func() {
// 		for d := range msgs {
// 			log.Printf("Received a message: %s", d.Body)
// 		}
// 	}()

// 	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
// 	<-forever
// 	return nil
// }

// func (s *ServerStorage) Send(config Config) error {
// 	conn, err := amqp.Dial(config.Address)
// 	if err != nil {
// 		return err
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		return err
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"hello", // name
// 		false,   // durable
// 		false,   // delete when unused
// 		false,   // exclusive
// 		false,   // no-wait
// 		nil,     // arguments
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	body := "Hello World!"
// 	err = ch.PublishWithContext(ctx,
// 		"",     // exchange
// 		q.Name, // routing key
// 		false,  // mandatory
// 		false,  // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		})

// 	if err != nil {
// 		return err
// 	}
// 	log.Printf(" [x] Sent %s\n", body)
// 	return nil
// }
