package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func Ping(URL string, timeout time.Duration) int {
	var client = http.Client{
		Timeout: timeout * time.Second,
	}
	request, err := client.Head(URL)
	if err != nil {
		var dnsError *net.DNSError
		if errors.As(err, &dnsError) {
			log.Println("DNS Error", dnsError)
		} else if os.IsTimeout(err) {
			log.Println("Timeout Error", err)
		} else {
			log.Println("Error", err)
		}
		return 0
	}
	defer request.Body.Close()
	return request.StatusCode
}
