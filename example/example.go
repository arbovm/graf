package main

import (
	"github.com/arbovm/graf"
	"log"
	"time"
	"math/rand"

)


func main() {
	g, err := graf.Dial("0.0.0.0:2003", 3*time.Second)
	if err != nil {
		log.Fatalf("Dialing graphite host failed. Error: %s", err.Error())
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		g.Send("test.example.random", r.Int63n(100), time.Now())
		if err != nil {
			log.Fatalf("Sending metrics to graphite host failed. Error: %s", err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
