package main

import (
  "log"

  amqp "github.com/rabbitmq/amqp091-go"
)

// Our receive has the same function to help us in error handling
func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

func main(){

  // My setup in Receive is the same as the publisher. We open a connection and channel, and declare the queue. The queue that sends publications

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

  // We declare the queue here, because we might start the consumer before the publisher. We quanto to make sure the queue exists before we try to consume mensagem from it.

  // Now, we are tell the server to delive us the message from queue
  msgs, err := ch.Consume(
  q.Name, // queue
  "",     // consumer
  true,   // auto-ack
  false,  // exclusive
  false,  // no-local
  false,  // no-wait
  nil,    // args
)
failOnError(err, "Failed to register a consumer")

var forever chan struct{}

// Using goroutine to push us messages asynchronously. So, we will read the messages from a channel

go func ()  {
  for d := range msgs{
    log.Printf("Received a message: %s", d.Body)
  }
}()

log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever

}

