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

	otr := NewOmdbTsvReader(WithContext(ctx), WithGoroutineCount(2), Filename("title.basics.tsv"))

	outputCh, err := otr.ReadAsync()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("IMDB_ID | Title | Plot")
	for omdbTitleRecord := range outputCh {
		if !omdbTitleRecord.applyFilters(programOptions) {
			continue
		}

		go func(omdbTitleRecord OmdbTitleRecord, ctx context.Context) {
			or, err := SearchForPlot(omdbTitleRecord, ctx)
			if err != nil {
				log.Printf("Got error while fetching OMDB record. Error: %e\n", err)
				return
			}

			fmt.Printf("%s | %s | %s\n", or.ImdbID, or.Title, or.Plot)
		}(omdbTitleRecord, ctx)
	}
}

func listenShutdown(cancel func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
}

func parseFlags() ProgramOptions {
	filePathFlag := flag.String("filePath", "", "Absolute path to the inflated `title.basics.tsv.gz` file")
	titleTypeFlag := flag.String("titleType", "", "Filter on `titleType` column")
	primaryTitleFlag := flag.String("primaryTitle", "", "Filter on `primaryTitle` column")
	originalTitleFlag := flag.String("originalTitle", "", "Filter on `originalTitle` column")
	genreFlag := flag.String("genre", "", "Filter on `genre` column")
	startYearFlag := flag.String("startYear", "", "Filter on `startYear` column")
	endYearFlag := flag.String("endYear", "", "Filter on `endYear` column")
	runtimeMinutesFlag := flag.String("runtimeMinutes", "", "Filter on `runtimeMinutes` column")
	genresFlag := flag.String("genres", "", "Filter on `genres` column")
	maxApiRequestsFlag := flag.Int("maxApiRequests", 10, "Maximum number of requests to be made to [omdbapi](https://www.omdbapi.com/)")
	maxRunTimeFlag := flag.Duration("maxRunTime", time.Second*10, "Maximum run time of the application. Format is a `time.Duration` string see [here](https://godoc.org/time#ParseDuration)")
	maxRequestsFlag := flag.Int("maxRequests", 10, "Maximum number of requests to send to [omdbapi](https://www.omdbapi.com/)")
	plotFilterFlag := flag.String("plotFilter", "", "Regex pattern to apply to the plot of a film retrieved from [omdbapi](https://www.omdbapi.com/)")

	flag.Parse()

	return ProgramOptions{
		filePathFlag:       *filePathFlag,
		titleTypeFlag:      *titleTypeFlag,
		primaryTitleFlag:   *primaryTitleFlag,
		originalTitleFlag:  *originalTitleFlag,
		genreFlag:          *genreFlag,
		startYearFlag:      *startYearFlag,
		endYearFlag:        *endYearFlag,
		runtimeMinutesFlag: *runtimeMinutesFlag,
		genresFlag:         *genresFlag,
		maxApiRequestsFlag: *maxApiRequestsFlag,
		maxRunTimeFlag:     *maxRunTimeFlag,
		maxRequestsFlag:    *maxRequestsFlag,
		plotFilterFlag:     *plotFilterFlag,
	}
}
