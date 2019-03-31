package structure

import (
	"time"
	"strings"
	"public-data/logger"
	"fmt"
	. "public-data/integrated/googlenews"
	. "public-data/integrated/fintimes"
	"github.com/joho/godotenv"
)

// Pub is a struct for holding Public data source members and base functions
type Pub struct {
	name string
}

func (p Pub) Name() string {
	return p.name
}

// interface function of base class
func (p Pub) GetMyTrades(pair Pair, start, end time.Time) TradeLogS {
	fmt.Println(p.name + ".GetMyTrades(pair Pair, start, end time.Time)" + " not implemented")
	return TradeLogS{}
}


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
