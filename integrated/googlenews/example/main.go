package main

import (
	"public-data/integrated/googlenews"
	"public-data/structure"
	"flag"
	"time"
)

func main()  {
	dbinfo := structure.DBinfo{structure.Address, structure.Port, structure.DBName}
	var dur time.Duration
	for {
		google := googlenews.NewGoogle()
		flag.DurationVar(&dur, "duration", 24, "query frequency (in hour)")
		flag.Parse()
		google.SetFreq(dur)
		raw := google.GetData()
		google.SaveDB(dbinfo, raw)
		time.Sleep(24*time.Hour)
	}
}
