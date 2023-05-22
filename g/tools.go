package g

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}

func IsTimeFormat(datetime string) bool {

	//m := "^([0-9]{4})(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])(2[0-3]|[01][0-9])([0-5][0-9])([0-5][0-9])$"
	m := "^([0-9]{4})(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])$"
	re := regexp.MustCompile(m)
	return re.MatchString(datetime)
}

func CheckAndCreateDir(fspath string) bool {

	// fspath ==  filePath

	dirName := filepath.Dir(fspath)
	if _, err := os.Stat(dirName); err != nil {
		_ = os.MkdirAll(dirName, 0755)
	}
	return true
}


func FillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}