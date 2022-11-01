package main

import "time"

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

type ProgramOption struct {
	PrimaryTitleFlag string
	MaxRuntime       time.Duration
}
