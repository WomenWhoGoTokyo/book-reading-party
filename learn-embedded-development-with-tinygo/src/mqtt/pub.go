package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// ClientIDのためのランダム値を生成する
	var rndbuf [4]byte
	rand.Read(rndbuf[:])
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(fmt.Sprintf(clientfmt, rndbuf))
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer c.Disconnect(250)
	for i := 1; ; i++ {
		// MQTTにPublish(送信)する
		text := fmt.Sprintf("mqtt from micchie %X %d", rndbuf, i)
		fmt.Printf("send : %s\n", text)
		token := c.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(5 * time.Second)
	}
}
