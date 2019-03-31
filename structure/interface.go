package structure

import "time"

// the supported public data sources
// NOTE: the names should all be uppercase
const (
	// add new data source here:)
	GoogleNews	= "GOOGLENEWS"
	FinTime     = "FINTIMES"
	// DB info
	Address		= "http://localhost"
	Port		= "8086"
	DBName		= "PUBLIC"
)

// Public is the interface for all pubs
// Any pub struct should implement all these interfaces
type Public interface {
	Name() string
	GetData() Datapoints
	SetFreq(dur time.Duration)
	SaveDB(dbinfo DBinfo, da Datapoints)
}


// return data structure
type Datapoint struct {
	Time     		time.Time
	Author	 		string
	Title    		string
	Description		string
	Body     		string
	Category 		string
	Source   		string
	URL  			string
}

type Datapoints []Datapoint

// database details
type DBinfo struct {
	IPAddr		string
	Port		string
	DBname		string
}

