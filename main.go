package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	programOptions := ParseFlags()

	ctx, cancel := context.WithTimeout(context.Background(), programOptions.maxRunTimeFlag)
	defer cancel()

	go listenShutdown(cancel)

	otr := NewOmdbTsvReader(WithContext(ctx), WithGoroutineCount(2), Filename(programOptions.filePathFlag))

	outputCh, err := otr.ReadAsync()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("IMDB_ID | Title | Plot")

	var wg sync.WaitGroup
	for omdbTitleRecord := range outputCh {
		if !omdbTitleRecord.applyFilters(programOptions) {
			continue
		}

		wg.Add(1)
		go func(omdbTitleRecord OmdbTitleRecord, ctx context.Context) {
			defer wg.Done()
			or, err := SearchForPlot(omdbTitleRecord, ctx)
			if err != nil {
				log.Printf("Got error while fetching OMDB record. Error: %e\n", err)
				return
			}

			fmt.Printf("%s | %s | %s\n", or.ImdbID, or.Title, or.Plot)
		}(omdbTitleRecord, ctx)
	}

	wg.Wait()
}

// listenShutdown runs invokes cancel method for graceful shutdown
func listenShutdown(cancel func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
}
