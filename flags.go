package main

import (
	"flag"
	"time"
)

func ParseFlags() ProgramOptions {
	filePathFlag := flag.String("filePath", "title.basics.tsv", "Absolute path to the inflated `title.basics.tsv.gz` file")
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
