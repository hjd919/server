package database

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
)

type ESConfig struct {
	Addresses []string `mapstructure:"dsn"`
}

func NewESConn(c *ESConfig) (es *elasticsearch.Client, err error) {
	log.Println("初始化es:")
	cfg := elasticsearch.Config{
		Addresses: c.Addresses,
		// Transport: &http.Transport{
		// 	MaxIdleConnsPerHost:   10,
		// 	ResponseHeaderTimeout: time.Millisecond,
		// 	DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
		// 	TLSClientConfig: &tls.Config{
		// 		MinVersion: tls.VersionTLS11,
		// 		// ...
		// 	},
		// },
		// Logger: &estransport.TextLogger{Output: os.Stdout},
		Logger: &estransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	}
	es, err = elasticsearch.NewClient(cfg)
	return es, err
}
