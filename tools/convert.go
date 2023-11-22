package tools

import (
	"time"
)

/* Convert string to float64 */
// This func is build to convert buy and sell values from scraped web data,
// that read as string.
func ConvertToFloat(s string) float64 {
	res, decimal := 0., 10.
	comma := false

	for i := range s {
		numb := float64(s[i] - '0')
		if s[i] == '.' {
			continue
		}
		if s[i] == ',' {
			comma = true
			continue
		}
		if comma == true {
			numb /= decimal
			decimal *= 10
		} else {
			res *= 10
		}
		res += numb
	}
	return res
}

// convert go time.Now() into date only
func GetDateOnly(input string) (time.Time, error) {
	//YYYYMMDD := "2022-01-20"
	t, err := time.Parse(time.DateOnly, input)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
