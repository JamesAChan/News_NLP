# News_NLP

## IMF Innovation Lab Unit: NLP Project

### Features
This repo is created for IMF project and aiming to realize the following features:
* Data acquisition: include historical data and real-time data
* Data storage: store different data types into different databases (support InfluxDB so far)
* Real-time monitor and alert: add slack and email notification channels to monitor server status and different processes
------------------
### Advantages
* Extendability: can add other data sources easily
* Speed: take the advantage of Golang and achieve fast data query/storage/processing
* Flexibility: can add new features into `structure/interface.go` by adding another functions:
```go
// Public is the interface for all pubs
// Any pub struct should implement all these interfaces
type Public interface {
	Name() string
	GetData() Datapoints
	SetFreq(dur time.Duration)
	SaveDB(dbinfo DBinfo, da Datapoints)
	//e.g. add new feature: conduct data cleaning 
	DataClean()
}
```
### Tips for usage (to be continue)
* add new data source into `const`:
```go
// the supported public data sources
// NOTE: the names should all be uppercase
const (
  // add new data source here:) 
	GoogleNews	= "GOOGLENEWS"
	FinTime     = "FINTIMES"
	//DB info
	Address		= "http://localhost"
	Port		= "8086"
	DBName		= "PUBLIC"
)
```
* Then add new data source structure into `structure.go`:
```go
// NewPub return an Public interface given an Public name
func NewPub(name string) Public {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	switch strings.ToUpper(name) {
	case GoogleNews:
		return NewGoogle()
	case FinTime:
		return NewFinTimes()
	default:
		logger.Panic().Str("pchagne name", name).Msg("invalid Pub")
		return nil
	}
}
```
* Next, implement all methods in `interface.go` for new data source
1. first implement basic functions in `implement` folder
2. secondly, integrate the basic functions with predefined methods
3. thirdly, write a `main.go` file for execution
4. lastly, add the execution file (the above `main.go` file) into `Makefile` to generate `binary file`, which will be located in `bin` folder

### Future features
* welcome to make any comments and give your precious suggestion for improvements




