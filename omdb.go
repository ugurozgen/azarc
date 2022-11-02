package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	OmdbApiUrl = "http://www.omdbapi.com"
	OmdbApiKey = "b020b1bb"
)

func SearchForPlot(otr OmdbTitleRecord, ctx context.Context) (*OmdbRecord, error) {
	requestURL := fmt.Sprintf("%s?i=%s&apikey=%s", OmdbApiUrl, otr.Tconst, OmdbApiKey)
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
