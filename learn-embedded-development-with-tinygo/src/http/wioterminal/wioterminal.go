package main

import (
	"fmt"
	"io"
	"log"
	"machine"
	"net/http"
	"net/url"
	"strings"
	"time"

	"tinygo.org/x/drivers/netlink"
	"tinygo.org/x/drivers/netlink/probe"
)

var (
	ssid string
	pass string
)

func main() {

	waitSerial()

	link, _ := probe.Probe()

	err := link.NetConnect(&netlink.ConnectParams{
		Ssid:       ssid,
		Passphrase: pass,
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "John Doe"
	occupation := "gardener"

	params := "name=" + url.QueryEscape(name) + "&" +
		"occupation=" + url.QueryEscape(occupation)

	path := fmt.Sprintf("https://httpbin.org/get?%s", params)

	cnt := 0
	for {
		fmt.Printf("Getting %s\r\n\r\n", path)
		resp, err := http.Get(path)
		if err != nil {
			fmt.Printf("%s\r\n", err.Error())
			time.Sleep(10 * time.Second)
			continue
		}

		fmt.Printf("%s %s\r\n", resp.Proto, resp.Status)
		for k, v := range resp.Header {
			fmt.Printf("%s: %s\r\n", k, strings.Join(v, " "))
		}
		fmt.Printf("\r\n")

		body, err := io.ReadAll(resp.Body)
		println(string(body))
		resp.Body.Close()

		cnt++
		fmt.Printf("-------- %d --------\r\n", cnt)
		time.Sleep(10 * time.Second)
	}
}

// Wait for user to open serial console
func waitSerial() {
	for !machine.Serial.DTR() {
		time.Sleep(100 * time.Millisecond)
	}
}
