package utlog

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticLogObject struct {
	client *elasticsearch.Client
}

func NewElasticLogClient(domain, user, password string, port int, isSecure bool) (eObj IElastic, err error) {
	protocol := "https"
	if !isSecure {
		protocol = "http"
	}

	if user == "" || password == "" || domain == "" {
		err = errors.New("user or password or domain should not empty")
		return
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("%s://%s:%d", protocol, domain, port),
		},
		Username: user,
		Password: password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Only for self-signed certificates, otherwise use a valid cert
			},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		Errorf("Error creating the client: %s", err)
		return
	}

	eObj = &ElasticLogObject{
		client: es,
	}

	return eObj, nil
}

func (e *ElasticLogObject) SendLog(appName string, message interface{}) error {
	jsonBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshaling log message: %v", err)
	}

	req := esapi.IndexRequest{
		Index:   appName,
		Body:    strings.NewReader(string(jsonBody)),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("error sending log message: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	// log.Println("Log message sent to Elasticsearch:", logMessage)
	return nil
}
