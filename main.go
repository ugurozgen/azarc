package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go listenShutdown(cancel)

	otr := NewOmdbTsvReader(WithContext(ctx), WithGoroutineCount(2))

	outputCh, err := otr.ReadAsync()
	if err != nil {
		log.Fatalln(err)
	}

	for title := range outputCh {
		fmt.Println(title)
	}
}

func listenShutdown(cancel func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
}
