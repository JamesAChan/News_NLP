package googlenews

import (
	"time"
	"os"
	"log"
	"fmt"
	. "public-data/structure"
	"public-data/implement/google"
	"github.com/joho/godotenv"
	"github.com/influxdata/influxdb/client/v2"
)

type Google struct {
	*Pub
	key  string
	Freq time.Duration
}

// constructor
func NewGoogle() Google {
	errL := godotenv.Load()
	if errL != nil {
		panic("Error loading .env file")
	}
	api := os.Getenv("API_GOOGLENEWS")
	return Google{Pub: &Pub{GoogleNews}, key:api, Freq:24}
}

// implement methods
func (p Google) Name() string {
	return p.Pub.Name()
}

func (p Google) SetFreq(dur time.Duration) {
	p.Freq = dur
	fmt.Println("Frequency set as:", dur, " Hour(s)")
	return
}

func (p Google) GetData() Datapoints {
	raw := google.Query(p.Freq, p.key)
	var datapts Datapoints
	var datapt  Datapoint
	for _, v := range raw {
		datapt.Time, _ = time.Parse("2006-01-02T15:04:05.000Z", v.Time)
		datapt.Author = v.Author
		datapt.Title = v.Title
		datapt.Description = v.Description
		datapt.Body = v.Content
		datapt.Category = v.Category
		datapt.Country = v.Country
		datapt.Source = v.Sources.Name
		datapt.URL = v.URL
		datapts = append(datapts, datapt)
	}
	return datapts
}

func (p Google) SaveDB(dbinfo DBinfo, da Datapoints) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     dbinfo.IPAddr + ":" + dbinfo.Port,
		//Username: username,
		//Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbinfo.DBname,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	for _, v := range da {
		fields := make(map[string]interface{})
		fields["Author"] = v.Author
		fields["Title"] = v.Title
		fields["Description"] = v.Description
		fields["Body"] = v.Body
		fields["Category"] = v.Category
		fields["Source"] = v.Source
		fields["URL"] = v.URL

		tags := map[string]string{
			"Source": v.Source,
			"Category": v.Category,
			"Country": v.Country,
		}
		pt, err := client.NewPoint(p.Name(), tags, fields, v.Time)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Name(), "saving finished at ", time.Now(), ":)")
}
