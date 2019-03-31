package db

import (
	"github.com/influxdata/influxdb/client/v2"
	"fmt"
	"log"
	"time"
)

const (
	MyDB = ""
	username = ""
	password = ""
)

// queryDB convenience function to query the database
func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}

	// Create DB
	_, err = queryDB(clnt, fmt.Sprintf("CREATE DATABASE %s", MyDB))
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func WriteUDP() {
	// Make client
	// PayloadSize is default value(512)
	c, err := client.NewUDPClient(client.UDPConfig{Addr: "localhost:8089"})
	if err != nil {
		panic(err.Error())
	}

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		panic(err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}


