package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/AsrofunNiam/lets-code-elastic-search/helper"
	"github.com/olivere/elastic/v7"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("https://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth("elastic", "you-credential"), // credential set
		elastic.SetHttpClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // Set disable TLS verification to self-signed
				},
			},
		}),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// search all documents
	// helper.SearchAllDocuments(client)

	// search by name and occupation
	helper.SearchByNameAndOccupation(client, "John Doe", "Developer")
}
