package helper

import (
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

func ApplyFilterElastic(client *elastic.Client, indexName string, filters *map[string]string) (*elastic.SearchService, error) {
	// search service
	searchService := client.Search().Index(indexName)
	query := elastic.NewBoolQuery()

	for key, value := range *filters {

		value = strings.ToLower(value)
		keyParts := strings.Split(key, ".")
		fieldName := keyParts[0]
		operator := keyParts[1]

		// Conditions filtered
		switch operator {
		case "eq":
			query = query.Must(elastic.NewTermQuery(fieldName, value))
		case "like":
			query = query.Must(elastic.NewRegexpQuery(fieldName, ".*"+value+".*"))
			// query = query.Must(elastic.NewWildcardQuery(fieldName, "*"+value+"*"))
		case "lt":
			query = query.Must(elastic.NewRangeQuery(fieldName).Lt(value))
		case "lte":
			query = query.Must(elastic.NewRangeQuery(fieldName).Lte(value))
		case "gt":
			query = query.Must(elastic.NewRangeQuery(fieldName).Gt(value))
		case "gte":
			query = query.Must(elastic.NewRangeQuery(fieldName).Gte(value))
		case "in":
			// convert value array to query terms
			values := strings.Split(value, ",")
			query = query.Must(elastic.NewTermsQuery(fieldName, values))
		default:
			return nil, fmt.Errorf("invalid operator: %s", operator)
		}
	}

	searchService = searchService.Query(query)
	return searchService, nil
}
