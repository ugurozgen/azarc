package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	OmdbApiUrl = "http://www.omdbapi.com"
)

// SearchForPlot fetches omdb record with title id
func SearchForPlot(otr OmdbTitleRecord, ctx context.Context) (*OmdbRecord, error) {
	requestURL := fmt.Sprintf("%s?i=%s&apikey=%s", OmdbApiUrl, otr.Tconst, os.Getenv("OMDB_API_KEY"))
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	var or OmdbRecord
	err = json.Unmarshal(body, &or)
	if err != nil {
		return nil, err
	}

	return &or, nil
}
