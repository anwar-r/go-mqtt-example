package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_publisher1")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting to broker: %v\n", token.Error())
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("test/topic", 0, false, text)
		token.Wait()
		time.Sleep(1 * time.Second)
	}

	client.Disconnect(250)
	fmt.Println("Publisher disconnected")
}
