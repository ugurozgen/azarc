package main

import (
	"context"
	"sync"
	"time"
)

type TsvReader interface {
	ReadAsync() error
}

type OmdbTsvReader struct {
	goroutineCount int
	fileName       string
	ctx            context.Context
	cancel         context.CancelFunc
	outputCh       chan OmdbTitleRecord
	wg             sync.WaitGroup
	recordCh       chan []string
}

type OmdbTitleRecord struct {
	Tconst         string
	TitleType      string
	PrimaryTitle   string
	OriginalTitle  string
	IsAdult        string
	StartYear      string
	EndYear        string
	RuntimeMinutes string
	Genres         string
}

type ProgramOptions struct {
	filePathFlag       string
	titleTypeFlag      string
	primaryTitleFlag   string
	originalTitleFlag  string
	genreFlag          string
	startYearFlag      string
	endYearFlag        string
	runtimeMinutesFlag string
	genresFlag         string
	maxApiRequestsFlag int
	maxRunTimeFlag     time.Duration
	maxRequestsFlag    int
	plotFilterFlag     string
}

type OmdbRecord struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}
