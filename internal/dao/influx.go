package dao

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Connector struct {
	Client influxdb2.Client
}

func (*Connector) init() Connector {
	token := "example-token"
	// Store the URL of your InfluxDB instance
	url := "https://us-west-2-1.aws.cloud2.influxdata.com"
	return Connector{Client: influxdb2.NewClient(url, token)}
}
