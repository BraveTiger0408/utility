package utility

import (
	"encoding/hex"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRoundFloat64(t *testing.T) {
	Convey("Round a float64 to x decimal places", t, func() {
		Convey("should return a valid rounded float64", func() {

			// define input and expected result
			type testdata struct {
				number float64
				places int
				result float64
			}

			// define test data
			var tests = []testdata{
				{27865428945.7825346789253, 2, 27865428945.78},
				{27865428945.7825346789253, 6, 27865428945.782535},
				{27865428945.7825346789253, 1, 27865428945.8},
				{0.7825346789253, 4, 0.7825},
				{0.782549000, 4, 0.7825},
				{-0.78255000, 4, -0.7826},
				{0.0, 0, 0},
			}

			// run tests
			for _, test := range tests {
				result := RoundFloat64(test.number, test.places)
				fmt.Printf("\n       %f, %f ", result, test.result)
				match := Float64Equal(result, test.result)
				So(match, ShouldEqual, true)
			}
		})
	})
}

func TestRoundDuration(t *testing.T) {
	Convey("Round a time.Duration", t, func() {
		Convey("should return a valid rounded duration", func() {

			// define input and expected result
			type testdata struct {
				period  time.Duration
				roundTo time.Duration
				result  time.Duration
			}

			// define test data
			var tests = []testdata{
				{500 * time.Millisecond, time.Second, 1 * time.Second},
				{499 * time.Millisecond, time.Second, 0 * time.Second},
				{500 * time.Millisecond, 0 * time.Second, 500 * time.Millisecond},
				{-500 * time.Millisecond, time.Second, -1 * time.Second},
			}

			// run tests
			for _, test := range tests {
				result := RoundDuration(test.period, test.roundTo)
				fmt.Printf("\n       %v, %v ", result, test.result)
				So(result, ShouldEqual, test.result)
			}
		})
	})
}

func TestFloat64Equal(t *testing.T) {
	Convey("compare float64 values", t, func() {
		Convey("should return equivalent `true` if values match", func() {

			// define input and expected result
			type testdata struct {
				valueOne float64
				valueTwo float64
				result   bool
			}

			// NOTE epsilon = 0.00000001

			// define test data
			var tests = []testdata{
				{0.0, 0.0, true},
				{0.0, 0.1, false},
				{0.0, 0.000000009, true},
				{0.0, 0.00000001, false}, //epsilon
				{0.0, 0.00000099, false},
				{435.9, 435.89999999, true},
			}

			// run tests
			for _, test := range tests {
				result := Float64Equal(test.valueOne, test.valueTwo)
				fmt.Printf("\n       %v, %v ", result, test.result)
				So(result, ShouldEqual, test.result)
			}
		})
	})
}

func TestMakeFileName(t *testing.T) {
	Convey("Making a new filename", t, func() {
		Convey("should return a unique filename with extension .ach", func() {
			filename := MakeFileName(".ach", "test/")
			ext := filepath.Ext(filename)
			So(filename, ShouldNotEqual, "")
			So(ext, ShouldEqual, ".ach")
		})
	})
}

func TestGetLocalIP(t *testing.T) {
	Convey("Function GetLocalIP", t, func() {
		Convey("Should return the local IP address", func() {
			result, err := GetLocalIP()
			So(result, ShouldNotEqual, "")
			So(err, ShouldEqual, nil)
		})
	})
}

func TestPadding(t *testing.T) {

	Convey("if string to be padded is longer than length (l)", t, func() {
		Convey("it should return an error", func() {
			res, err := Padding("abc", 1, "right", " ")
			So(res, ShouldEqual, "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "string is too long")
		})
	})

	Convey("if justified is not \"right\" or \"left\"", t, func() {
		Convey("it should return an error", func() {
			res, err := Padding("abc", 10, "hello", " ")
			So(res, ShouldEqual, "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "justification must be either right or left")
		})
	})

	Convey("if Padding character is longer than 1 character", t, func() {
		Convey("it should return an error", func() {
			res, err := Padding("abc", 10, "right", "12")
			So(res, ShouldEqual, "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "padding must be only one character")
		})
	})

	Convey("if string is \"abc\" and length is 10, and justification is \"left\"", t, func() {
		Convey("it should return \"abc       \"", func() {
			res, err := Padding("abc", 10, "left", " ")
			So(res, ShouldEqual, "abc       ")
			So(err, ShouldBeNil)
		})
	})

	Convey("if string is \"abc\" and length is 10, and justification is \"right\"", t, func() {
		Convey("it should return \"       abc\"", func() {
			res, err := Padding("abc", 10, "right", " ")
			So(res, ShouldEqual, "       abc")
			So(err, ShouldBeNil)
		})
	})

	Convey("if string is \"abc\", length is 10, justification is \"right\" and Padding character is \"0\"", t, func() {
		Convey("it should return \"0000000abc\"", func() {
			res, err := Padding("abc", 10, "right", "0")
			So(res, ShouldEqual, "0000000abc")
			So(res, ShouldNotEqual, "       abc")
			So(err, ShouldBeNil)
		})
	})

}

func TestCalcSettlementDate(t *testing.T) {
	Convey("Calculate settement date", t, func() {
		Convey("should return a valid settlement date", func() {

			const timeFormat = "060102"

			// Make a map of bank holidays
			var bankHolidayMap = make(map[string]bool)

			var holidayList = []string{
				"160704",
				"160904",
				"161009",
				"161110",
			}

			for _, holiday := range holidayList {
				bankHolidayMap[holiday] = true
			}

			// define input and expected result
			type testpair struct {
				date   string
				result string
			}

			// define test data
			var tests = []testpair{
				{"160629", "160630"}, // regular weekday
				{"160630", "160701"}, // Thursday
				{"160701", "160705"}, // weekend & *holiday
				{"160708", "160711"}, // Friday
				{"160709", "160711"}, // Saturday
				{"160710", "160711"}, // Sunday
				{"160704", "160705"}, // *holiday
			}

			// run tests
			for _, testData := range tests {
				testDate, _ := time.Parse(timeFormat, testData.date)
				result := CalcSettlementDate(testDate, bankHolidayMap)
				So(result.Format(timeFormat), ShouldEqual, testData.result)
			}
		})
	})
}

func TestGenRandomString(t *testing.T) {
	Convey("Generate a hexadecimal random string", t, func() {
		Convey("should return a valid hexadecimal string of length l", func() {

			// define input and expected result
			type testdata struct {
				length int
				result int
			}

			// define test data.  Hexidecimal
			// is double the length
			var tests = []testdata{
				{0, 0}, // err test
				{1, 2},
				{2, 4},
				{20, 40},
				{-1, 0}, // err test
			}

			// run tests
			for _, test := range tests {
				result := GenRandomString(test.length)
				// make sure it's hex (we can decode it)
				_, err := hex.DecodeString(result)
				So(err, ShouldBeNil)
				So(len(result), ShouldEqual, test.result)
			}
		})
	})
}
