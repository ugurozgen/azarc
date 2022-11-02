package main

import (
	"flag"
	"strings"
)

type FilterFunc func(po ProgramOptions) bool

func (otr OmdbTitleRecord) applyFilters(po ProgramOptions) bool {
	filters := []FilterFunc{
		otr.primaryTitleFilter,
		otr.originalTitleFilter,
		otr.genreFilter,
		otr.startYearFilter,
		otr.endYearFilter,
		otr.runtimeMinutesFilter,
		otr.genresFilter,
	}

	valid := true
	for _, filter := range filters {
		if !filter(po) {
			valid = false
		}
	}

	return valid
}

func (otr OmdbTitleRecord) primaryTitleFilter(po ProgramOptions) bool {
	return !isFlagPassed("primaryTitle") ||
		(po.primaryTitleFlag != "" && strings.Contains(strings.ToLower(otr.PrimaryTitle), po.primaryTitleFlag))
}

func (otr OmdbTitleRecord) originalTitleFilter(po ProgramOptions) bool {
	return !isFlagPassed("originalTitle") ||
		(po.originalTitleFlag != "" && strings.Contains(strings.ToLower(otr.OriginalTitle), po.originalTitleFlag))
}

func (otr OmdbTitleRecord) genreFilter(po ProgramOptions) bool {
	return !isFlagPassed("genre") ||
		(po.genreFlag != "" && strings.Contains(strings.ToLower(otr.Genres), po.genreFlag))
}

func (otr OmdbTitleRecord) startYearFilter(po ProgramOptions) bool {
	return !isFlagPassed("startYear") ||
		(po.startYearFlag != "" && strings.Contains(strings.ToLower(otr.StartYear), po.startYearFlag))
}

func (otr OmdbTitleRecord) endYearFilter(po ProgramOptions) bool {
	return !isFlagPassed("endYear") ||
		(po.endYearFlag != "" && strings.Contains(strings.ToLower(otr.EndYear), po.endYearFlag))
}

func (otr OmdbTitleRecord) runtimeMinutesFilter(po ProgramOptions) bool {
	return !isFlagPassed("runtimeMinutes") ||
		(po.runtimeMinutesFlag != "" && strings.Contains(strings.ToLower(otr.RuntimeMinutes), po.runtimeMinutesFlag))
}

func (otr OmdbTitleRecord) genresFilter(po ProgramOptions) bool {
	return !isFlagPassed("genres") ||
		(po.genresFlag != "" && strings.Contains(strings.ToLower(otr.Genres), po.genresFlag))
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
