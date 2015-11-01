package main
import (
	"strconv"
	"fmt"
	//"google.golang.org/api/bigquery/v2"
	"net/http"
	"net/url"
	"encoding/json"
)

type Request struct {
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

type AppConf struct {
	ArtistName string `json:"artistName`
	IsGameCenterEnabled bool `json:"isGameCenterEnabled"`
	TrackId int `json:"trackId"`
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
	Limit: 200,
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

func main() {

    var results Request

    url := appStoreQuery.getUrl()
    fmt.Println(url)

	err := getJson(url, &results)
    if err != nil {
    	fmt.Println(err)
        return
    }
	fmt.Println(results)//.ResultCount)

	for _, value := range results.Results {
		fmt.Println(value.TrackId)
	}
    //bigqueryService, err := bigquery.New(oauthHttpClient)
}