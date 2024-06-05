package main

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_consumer")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting to broker: %v\n", token.Error())
		os.Exit(1)
	}

	if token := client.Subscribe("test/topic", 0, messageHandler); token.Wait() && token.Error() != nil {
		fmt.Printf("Error subscribing to topic: %v\n", token.Error())
		os.Exit(1)
	}

	fmt.Println("Subscribed to topic: test/topic")

	// Keep the consumer running
	select {}
}
