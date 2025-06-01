package main

import (
  "context"
  "log"
  "time"

  amqp "github.com/rabbitmq/amqp091-go"
)

// Helper Function to return value for each amqp call
func failOnError(err error , msg string){
    if err != nil {
      log.Panicf("%s: %s", msg, err)
    }
}

func main() {
    // Connect to RabbitMQ server
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to open a channel")
    defer conn.Close()

    /* The connection abstracts the socket connection, and takes care of protocol negotiation and auth. Its nice! Now, we create a channel to calls the most of the API things*/

    ch, err := conn.Channel()
    failOnError(err, "Failed to open the channel")
    defer ch.Close()

    // Now, we need to DeclareQueue for us to send
    q, err := ch.QueueDeclare(
    "hello", // name
    false,   // durable
    false,   // delete when unused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
  )
  failOnError(err, "Failed to declare a queue")

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  // This is the message to we can publish to the queue
  body := "Hello World!"
  err = ch.PublishWithContext(ctx,
    "",     // exchange
    q.Name, // routing key
    false,  // mandatory
    false,  // immediate
    amqp.Publishing {
      ContentType: "text/plain",
      Body:        []byte(body),
    })
  failOnError(err, "Failed to publish a message")
  log.Printf(" [x] Sent %s\n", body)

  // Its done, Our consumer listen to RabbitMQ messages and listen for messages and print then out
}