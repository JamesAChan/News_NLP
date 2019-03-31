package fintimes

import (
	. "public-data/structure"
)
type FinTimes struct {
	*Pub
}

// constructor
func NewFinTimes() FinTimes {
	return FinTimes{Pub: &Pub{FinTime}}
}

func (p FinTimes) GetRawData() {

}
func (p FinTimes) Name() string {
	return ""
}

func (p FinTimes) SaveDB() {
	panic("Google.PlaceLimitOrder not implemented")
}

