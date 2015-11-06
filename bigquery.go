package main

import (
	"log"
	"container/list"
	"net/http"
	//"fmt"
    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
	bigquery "google.golang.org/api/bigquery/v2"
	//storage "google.golang.org/api/storage/v1"

	//"encoding/json"
)

const (
	GB                         = 1 << 30
	MaxBackoff                 = 30000
	BaseBackoff                = 250
	BackoffGrowthFactor        = 1.8
	BackoffGrowthDamper        = 0.25
	JobStatusDone              = "DONE"
	DatasetAlreadyExists       = "Already Exists: Dataset"
	TableWriteEmptyDisposition = "WRITE_EMPTY"
)


// Wraps the BigQuery service and dataset and provides some helper functions.
type bqDataset struct {
	ProjectId string
	DatasetId string
	TableId string
	bq      *bigquery.Service
	dataset *bigquery.Dataset
	jobsets map[string]*list.List
}


func newBQDataset(client *http.Client, projectId string, datasetId string, tableId string) (*bqDataset,
	error) {

	service, err := bigquery.New(client)
	if err != nil {
		log.Fatalf("Unable to create BigQuery service: %v", err)
	}

	return &bqDataset{
		ProjectId: projectId,
		DatasetId: datasetId,
		TableId:   tableId,
		bq:      service,
		dataset: &bigquery.Dataset{
			DatasetReference: &bigquery.DatasetReference{
				DatasetId: datasetId,
				ProjectId: projectId,
			},
		},
		jobsets: make(map[string]*list.List),
	}, nil
}
/*
func (ds *bqDataset) insert(existsOK bool) error {
	call := ds.bq.Datasets.Insert(ds.ProjectId, ds.DatasetId)
	_, err := call.Do()
	if err != nil && (!existsOK || !strings.Contains(err.Error(),
		DatasetAlreadyExists)) {
		return err
	}

	return nil
}
*/

type Children struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Person struct {
	FullName string `json:"fullName"`
	Age int `json:"age"`
	Childrens []Children `json:"children"`
} 


func (ds *bqDataset) Insert(request *AppRequest) error {
  	rows := make([]*bigquery.TableDataInsertAllRequestRows, request.size())
 
    for i, app := range request.Results {
    	rows[i] = new(bigquery.TableDataInsertAllRequestRows)
    	Json, err := app.getJson()
    	if err != nil {
    		return err
    	}
    	rows[i].Json = Json
    }

	insertRequest := &bigquery.TableDataInsertAllRequest{Rows: rows}
	//fmt.Println(ds.ProjectId, ds.DatasetId, ds.TableId)
    _, err := ds.bq.Tabledata.InsertAll(ds.ProjectId, ds.DatasetId, ds.TableId, insertRequest).Do()
	return err
}


func connectBigQueryDB() (*bqDataset, error) {

	projectId := "cometiphrd"
	datasetId := "appstore"//"october"
	tableId := "data"//"data_test"

	// Use oauth2.NoContext if there isn't a good context to pass in.
    ctx := context.Background()
    ts, err := google.DefaultTokenSource(ctx, bigquery.BigqueryScope,
    										  //storage.DevstorageReadOnlyScope,
    										  "https://www.googleapis.com/auth/userinfo.profile")
	if err != nil {
	        //...
	}
	client := oauth2.NewClient(ctx, ts)

	return newBQDataset(client, projectId, datasetId, tableId)
}