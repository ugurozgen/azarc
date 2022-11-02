package main

import (
	"flag"
	"strings"
)

type FilterFunc func(po ProgramOptions) bool

func (otr OmdbTitleRecord) applyFilters(po ProgramOptions) bool {
	filters := []FilterFunc{}
	if isFlagPassed("primaryTitle") {
		filters = append(filters, otr.primaryTitleFilter)
	}
	if isFlagPassed("originalTitle") {
		filters = append(filters, otr.originalTitleFilter)
	}
	if isFlagPassed("genre") {
		filters = append(filters, otr.genreFilter)
	}
	if isFlagPassed("startYear") {
		filters = append(filters, otr.startYearFilter)
	}
	if isFlagPassed("endYear") {
		filters = append(filters, otr.endYearFilter)
	}
	if isFlagPassed("runtimeMinutes") {
		filters = append(filters, otr.runtimeMinutesFilter)
	}
	if isFlagPassed("genres") {
		filters = append(filters, otr.genresFilter)
	}

	valid := true
	for _, filter := range filters {
		if !filter(po) {
			valid = false
			break
		}
	}

	return valid
}

func (otr OmdbTitleRecord) primaryTitleFilter(po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.PrimaryTitle), strings.ToLower(po.primaryTitleFlag))
}

func (otr OmdbTitleRecord) originalTitleFilter(po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.OriginalTitle), strings.ToLower(po.originalTitleFlag))
}

func (otr OmdbTitleRecord) genreFilter(po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.Genres), strings.ToLower(po.genreFlag))
}

func (otr OmdbTitleRecord) startYearFilter(po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.StartYear), strings.ToLower(po.startYearFlag))
}

func (otr OmdbTitleRecord) endYearFilter(po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.EndYear), strings.ToLower(po.endYearFlag))
}

func (otr OmdbTitleRecord) runtimeMinutesFilter(po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.RuntimeMinutes), strings.ToLower(po.runtimeMinutesFlag))
}

func (otr OmdbTitleRecord) genresFilter(po ProgramOptions) bool {
	return strings.Contains(strings.ToLower(otr.Genres), strings.ToLower(po.genresFlag))
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
