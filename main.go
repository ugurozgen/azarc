package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	programOptions := parseFlags()

	ctx, cancel := context.WithTimeout(context.Background(), programOptions.MaxRuntime)
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

func parseFlags() *ProgramOption {
	primaryTitleFlag := flag.String("primaryTitle", "foo", "Primary Title filter that will be applied to data.")
	maxRunTimeFlag := flag.Duration("maxRunTime", time.Second*10, "Timeout for the operation. Default value is 10 seconds")

	flag.Parse()

	return &ProgramOption{
		PrimaryTitleFlag: *primaryTitleFlag,
		MaxRuntime:       *maxRunTimeFlag,
	}
}
