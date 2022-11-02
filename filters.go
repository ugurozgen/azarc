package main

import (
	"flag"
	"strings"
)

type FilterFunc func(po ProgramOptions) bool

func (otr OmdbTitleRecord) applyFilters(po ProgramOptions) bool {
	filters := []FilterFunc{
		otr.PrimaryTitleFilter,
	}

	valid := true
	for _, filter := range filters {
		if !filter(po) {
			valid = false
		}
	}

	return valid
}

func (otr OmdbTitleRecord) PrimaryTitleFilter(po ProgramOptions) bool {
	return isFlagPassed("primaryTitle") && strings.Contains(strings.ToLower(otr.PrimaryTitle), po.PrimaryTitleFlag)
}

func (otr OmdbTitleRecord) PrimaryTitleFilter(po ProgramOptions) bool {
	return isFlagPassed("primaryTitle") && strings.Contains(strings.ToLower(otr.PrimaryTitle), po.PrimaryTitleFlag)
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
