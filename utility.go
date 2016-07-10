package utility

import (
	"errors"
	"fmt"
	"math"
	"net"
	"path/filepath"
	"time"
)

const (
	timeFormat = "2006-01-02T15-04-05.000"
)

// var languages map[string]string
//
// func init() {
// 	languages = make(map[string]string)
// 	languages["cs"] = "C #"
// 	languages["js"] = "JavaScript"
// 	languages["rb"] = "Ruby"
// 	languages["go"] = "Golang"
// }
//
// func Get(key string) string {
// 	return languages[key]
// }
//
// func Add(key, value string) {
// 	languages[key] = value
// }
//
// func GetAll() map[string]string {
// 	return languages
// }

// RoundFloat64 rounds numbers
func RoundFloat64(val float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= float64(.5) {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

// RoundDuration rounds time.Duration
// Example to seconds: RoundDuration(d, time.Second)
func RoundDuration(d, r time.Duration) time.Duration {
	if r <= 0 {
		return d
	}
	neg := d < 0
	if neg {
		d = -d
	}
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	if neg {
		return -d
	}
	return d
}

// Float64Equal checks not whether the numbers are exactly the same, but
// whether their difference is very small. The error margin that the
// difference is compared to is called epsilon.
func Float64Equal(a, b float64) bool {
	const epsilon float64 = 0.00000001
	if (a-b) < epsilon && (b-a) < epsilon {
		return true
	}
	return false
}

// MakeFileName creates a new filename
func MakeFileName(fileExtension string, directory string) string {
	t := time.Now().UTC()
	timestamp := t.Format(timeFormat)
	return filepath.Join(directory, fmt.Sprintf("%s%s", timestamp, fileExtension))
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// check the address type (not a loopback)
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return ip, nil
			}
		}
	}
	return "", errors.New("no ip address found")
}

// Padding takes a string, length and either "left" or "right" for
// padding/justification plus a padding character.  It will pad the string out.
func Padding(s string, length int, justified string, padChar string) (string, error) {
	paddingNeeded := length - len(s)
	paddedString := ""

	// Check for errors
	if len(s) > length {
		return "", errors.New("string is too long")
	}
	if len(padChar) > 1 {
		return "", errors.New("padding must be only one character")
	}

	if justified == "right" {
		for i := 0; i < paddingNeeded; i++ {
			paddedString = paddedString + padChar
		}
		paddedString = paddedString + s
		return paddedString, nil
	}

	if justified == "left" {
		paddedString = s
		for i := 0; i < paddingNeeded; i++ {
			paddedString = paddedString + padChar
		}
		return paddedString, nil
	}

	return "", errors.New("justification must be either right or left")
}

// CalcSettlementDate takes in today's date (basically "now") along with a
// a map of bank holidays and it calculates to correct settlement date for
// ACH transactions.
func CalcSettlementDate(today time.Time, bankHolidayMap map[time.Time]bool) (settlementDate time.Time) {

	// NOTE: There is no definition in Go
	// for units of Day or larger to avoid confusion across
	// daylight savings time zone transitions. So we use hours.

	// settlement is tomorrow.
	settlementDate = today.Add(time.Hour * 24)

	// unless tomorrow is Saturday, then it's Monday
	if "Saturday" == settlementDate.Weekday().String() {
		// add two more days (cover the weekend)
		settlementDate = settlementDate.Add(time.Hour * 48)
	}

	// unless tomorrow is Sunday, then it's Monday
	if "Sunday" == settlementDate.Weekday().String() {
		// add one more day (cover the weekend)
		settlementDate = settlementDate.Add(time.Hour * 24)
	}

	// unless the calculated settlement day is a
	// holiday, then add one more day.
	_, found := bankHolidayMap[settlementDate]
	if found {
		settlementDate = settlementDate.Add(time.Hour * 24)
	}

	return settlementDate
}
