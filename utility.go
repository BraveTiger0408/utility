package utility

import (
	"errors"
	"fmt"
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
		return "", errors.New("String is too long")
	}
	if len(padChar) > 1 {
		return "", errors.New("Padding must be only one character")
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

	return "", errors.New("Justification must be either right or left")
}
