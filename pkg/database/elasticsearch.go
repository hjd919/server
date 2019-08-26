package database

// https://github.com/olivere/elastic/wiki

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"

	// "github.com/olivere/elastic"
	"github.com/olivere/elastic/v7"
)

// 接口
type ESClient interface {
	Connect()
}

type esClient struct {
	config *ESConfig
	Client *elastic.Client
}

type ESConfig struct {
	Addresses []string `mapstructure:"addresses"`
}

func (es *esClient) Connect() {
	// client, err := elastic.NewClient()
	// // client, err := elastic.NewClient(elastic.SetURL(es.config.Addresses...))
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// defer client.Stop()
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	es.Client = client
	return
}

// esClientPro
type esClientPro struct {
	config *ESConfig
	Client *elasticsearch.Client
}

func (esclientpro *esClientPro) Connect() {
	log.Println("初始化es:")
	cfg := elasticsearch.Config{
		Addresses: esclientpro.config.Addresses,
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
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		// Handle error
		panic(err)
	}
	esclientpro.Client = es
	return
}

func NewESClient(c *ESConfig) (client *elastic.Client) {
	// client, err := elastic.NewClient(elastic.SetURL(c.Addresses...))
	client, err := elastic.NewClient(elastic.SetURL(c.Addresses...), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return
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

func ESConn() (es *elastic.Client) {
	conf := &ESConfig{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es = NewESClient(conf)
	return es
}

func ESConnect() (es *elasticsearch.Client) {
	conf := &ESConfig{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, _ = NewESConn(conf)
	return es
}
