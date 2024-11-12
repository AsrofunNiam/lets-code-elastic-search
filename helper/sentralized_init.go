package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func LogToElasticsearch(message, level string, fields logrus.Fields, client *elastic.Client) error {
	if client == nil {
		return fmt.Errorf("Elasticsearch client not initialized")
	}

	logEntry := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"message":   message,
		"level":     level,
		"fields":    fields,
	}

	// Convert logEntry ke JSON
	data, err := json.Marshal(logEntry)
	if err != nil {
		return fmt.Errorf("error encoding log entry: %v", err)
	}

	// Save log ke index "logs" di Elasticsearch
	_, err = client.Index().
		Index("logs").
		BodyJson(string(data)).
		Do(context.Background())
	if err != nil {
		return fmt.Errorf("error indexing log: %v", err)
	}

	return nil
}
