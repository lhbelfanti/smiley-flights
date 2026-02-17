package search

import (
	"fmt"
	"net/url"
	"time"
)

type Parameters struct {
	Origin         string
	Destination    string
	DepartureDate  string
	ReturnDate     string
	Adults         int
	Children       int
	Infants        int
	IsFlexibleDate bool
	TripType       int
	CabinType      string
	CurrencyCode   string
}

func (f Parameters) ToQueryString() string {
	v := url.Values{}
	v.Add("originAirportCode", f.Origin)
	v.Add("destinationAirportCode", f.Destination)
	v.Add("departureDate", f.formatDate(f.DepartureDate))
	v.Add("returnDate", f.formatDate(f.ReturnDate))
	v.Add("adults", fmt.Sprintf("%d", f.Adults))
	v.Add("children", fmt.Sprintf("%d", f.Children))
	v.Add("infants", fmt.Sprintf("%d", f.Infants))
	v.Add("isFlexibleDateChecked", fmt.Sprintf("%t", f.IsFlexibleDate))
	v.Add("tripType", fmt.Sprintf("%d", f.TripType))
	v.Add("cabinType", f.CabinType)
	v.Add("currencyCode", f.CurrencyCode)
	return v.Encode()
}

func (f Parameters) formatDate(dateStr string) string {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "0"
	}
	return fmt.Sprintf("%d", t.UnixMilli())
}
