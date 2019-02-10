package ratesapi_go

import (
	"strings"
	"time"
)

type RatesDate time.Time

const ratesDateFormat = "2006-01-02"

func (ft *RatesDate) UnmarshalJSON(src []byte) (err error) {
	s := string(src)
	s = cleanString(s)
	t, err := time.Parse(ratesDateFormat, s)
	*ft = RatesDate(t)
	return
}

func cleanString(s string) string {
	return strings.Replace(s, `"`, "", -1)
}
