package main

import (
	"flag"
	"strings"
)

// FilterFunc represents filtering functions that come with program arguments
type FilterFunc func(otr OmdbTitleRecord, po ProgramOptions) bool

func parseFilters(po ProgramOptions) []FilterFunc {
	filters := []FilterFunc{}
	if isFlagPassed("primaryTitle") {
		filters = append(filters, primaryTitleFilter)
	}
	if isFlagPassed("originalTitle") {
		filters = append(filters, originalTitleFilter)
	}
	if isFlagPassed("genre") {
		filters = append(filters, genreFilter)
	}
	if isFlagPassed("startYear") {
		filters = append(filters, startYearFilter)
	}
	if isFlagPassed("endYear") {
		filters = append(filters, endYearFilter)
	}
	if isFlagPassed("runtimeMinutes") {
		filters = append(filters, runtimeMinutesFilter)
	}
	if isFlagPassed("genres") {
		filters = append(filters, genresFilter)
	}

	return filters
}

// applyFilters applies all filters that passed with arguments to program
func (otr OmdbTitleRecord) applyFilters(filters []FilterFunc, po ProgramOptions) bool {
	valid := true
	for _, filter := range filters {
		if !filter(otr, po) {
			valid = false
			break
		}
	}

	return valid
}

// primaryTitleFilter is a FilterFunc that checks primaryTitle argument
func primaryTitleFilter(otr OmdbTitleRecord, po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.PrimaryTitle), strings.ToLower(po.primaryTitleFlag))
}

// originalTitleFilter is a FilterFunc that check originalTitle argument
func originalTitleFilter(otr OmdbTitleRecord, po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.OriginalTitle), strings.ToLower(po.originalTitleFlag))
}

// genreFilter is a FilterFunc that check genre argument
func genreFilter(otr OmdbTitleRecord, po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.Genres), strings.ToLower(po.genreFlag))
}

// startYearFilter is a FilterFunc that check startYear argument
func startYearFilter(otr OmdbTitleRecord, po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.StartYear), strings.ToLower(po.startYearFlag))
}

// endYearFilter is a FilterFunc that check endYear argument
func endYearFilter(otr OmdbTitleRecord, po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.EndYear), strings.ToLower(po.endYearFlag))
}

// runtimeMinutesFilter is a FilterFunc that check runtimeMinutes argument
func runtimeMinutesFilter(otr OmdbTitleRecord, po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.RuntimeMinutes), strings.ToLower(po.runtimeMinutesFlag))
}

// genresFilter is a FilterFunc that check genres argument
func genresFilter(otr OmdbTitleRecord, po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.Genres), strings.ToLower(po.genresFlag))
}

// isFlagPassed checks if the argument passed with program call
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
