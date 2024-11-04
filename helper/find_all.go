package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

func SearchAllDocuments(client *elastic.Client) {
	// context
	ctx := context.Background()

	// Find All
	query := elastic.NewMatchAllQuery()

	// Execute
	searchResult, err := client.Search().
		Index("marketing-user").
		Query(query).
		Size(100).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error executing search query: %s", err)
	}

	fmt.Printf("Found %d documents\n", searchResult.TotalHits())

	// Iterate over results
	for _, hit := range searchResult.Hits.Hits {
		fmt.Printf("Document ID: %s\n", hit.Id)

		// Decode
		var doc map[string]interface{}
		err := json.Unmarshal(hit.Source, &doc)
		if err != nil {
			log.Printf("Error decoding document source: %s", err)
			continue
		}

		// Print JSON response
		docJSON, _ := json.MarshalIndent(doc, "", "  ")
		fmt.Printf("Source: %s\n", docJSON)
	}
}