package app

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AsrofunNiam/lets-code-elastic-search/model/domain"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	//  auto migrate
	err = database.AutoMigrate(
		domain.Product{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}

	return database
}

func ConnectionElastic(host, port, user, password string) *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL("http://"+host+":"+port),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth(user, password),
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
	return client
}
