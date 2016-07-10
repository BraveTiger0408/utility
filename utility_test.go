package utility

import (
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
			So(err.Error(), ShouldEqual, "String is too long")
		})
	})

	Convey("if justified is not \"right\" or \"left\"", t, func() {
		Convey("it should return an error", func() {
			res, err := Padding("abc", 10, "hello", " ")
			So(res, ShouldEqual, "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Justification must be either right or left")
		})
	})

	Convey("if Padding character is longer than 1 character", t, func() {
		Convey("it should return an error", func() {
			res, err := Padding("abc", 10, "right", "12")
			So(res, ShouldEqual, "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Padding must be only one character")
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
