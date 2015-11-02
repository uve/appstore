package main

import (
	"strconv"
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
)

type AppRequest struct {
    ResultCount int `json:"resultCount"`
    Results  []App  `json:"results"`
}

func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

// http://www.apple.com/itunes/affiliates/
// resources/documentation/itunes-store-web-service-search-api.html

type AppStoreQuery struct {
	BaseUrl string
	Limit int
	Country string
	Lang string
	Entity string
	Term string
}

var appStoreQuery = AppStoreQuery{
	BaseUrl: "https://itunes.apple.com/search?",
	Limit: 25,
	Country: "us",
	Lang: "en_us",
	Entity: "software",
	Term: "x",
}

func (query *AppStoreQuery) getUrl() string {	
	params := url.Values{}
	params.Add("entity",  query.Entity)
	params.Add("country", query.Country)
	params.Add("lang", query.Lang)
	params.Add("term", query.Term)
	params.Add("limit", strconv.Itoa(query.Limit))

	return query.BaseUrl + params.Encode()
}

func (request *AppRequest) find() {
	url := appStoreQuery.getUrl()
    //fmt.Println(url)

	err := getJson(url, &request)
    if err != nil {
    	fmt.Println(err)
        return
    }
	//fmt.Println(results.ResultCount)
}

func (request *AppRequest) filter(tracks []string) {
    uniq := make([]App, 0)
    for _, app := range request.Results {
        if !InSlice(app.TrackId, tracks) {
            uniq = append(uniq, app)
        }
    }
    request.Results = uniq
}

func (request *AppRequest) getTrackIds() []string {
    items := make([]string, request.size())
    for i, app := range request.Results {
        items[i] = strconv.Itoa(app.TrackId)
    }
    return items
}

func (request *AppRequest) save() {
	fmt.Println("Saved apps: ", request.size())
}

func (request *AppRequest) size() int {
	return len(request.Results)
}