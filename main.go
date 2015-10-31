package main
import (
	//"strconv"
	"fmt"
	//"google.golang.org/api/bigquery/v2"
	"net/http"
	"encoding/json"
)

type App struct {
	ArtistName string `json:"artistName`
	IsGameCenterEnabled bool `json:"isGameCenterEnabled"`
	TrackId int `json:"trackId"`
}

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

func main() {

    var results Request

    url := "https://itunes.apple.com/search?term=xxx&country=us&entity=software"
	err := getJson(url, &results)
    if err != nil {
    	fmt.Println(err)
        return
    }
	fmt.Println(results)
    //bigqueryService, err := bigquery.New(oauthHttpClient)
}