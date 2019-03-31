package structure

import (
	"strings"
	"public-data/logger"
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
