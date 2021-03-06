package logarchive

import (
	"os"
	"strconv"
	"strings"
	"time"
	"path/filepath"
	"errors"
	"fmt"
)

func GetExistsAbsolutePath(p string) (string, error) {
	v, err := filepath.Abs(p)
	if err != nil {
		return "", err;
	}

	f, err := os.Stat(v)
	if err == nil && f.IsDir() {
		return v, nil
	}

	if os.IsNotExist(err) {
		return "", err
	}

	return v, nil
}

func Mkdirp(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

func MkdirpByFileName(fileName string) error {
	pathSeparator := string(os.PathSeparator)
	s := strings.Split(fileName, pathSeparator)
	s2 := s[:len(s)-1]

	dir := strings.Join(s2, pathSeparator)

	return Mkdirp(dir)
}

func LeftPadInt(n int, v int) string {
	s := strconv.Itoa(v)
	for i := len(s); i < n; i++ {
		s = "0" + s
	}

	return s
}

type TimeInfo struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
	Time   *time.Time
}

func GetTime() TimeInfo {
	t := time.Now()
	d := TimeInfo{
		Year:   LeftPadInt(4, t.Year()),
		Month:  LeftPadInt(2, int(t.Month())),
		Day:    LeftPadInt(2, t.Day()),
		Hour:   LeftPadInt(2, t.Hour()),
		Minute: LeftPadInt(2, t.Minute()),
		Second: LeftPadInt(2, t.Second()),
		Time:   &t,
	}

	return d
}

func (t TimeInfo) Format(f string) string {
	r := []rune(f)
	var s string
	for _, v := range r {
		switch v {
		case 'Y':
			s += t.Year
		case 'm':
			s += t.Month
		case 'd':
			s += t.Day
		case 'H':
			s += t.Hour
		case 'i':
			s += t.Minute
		case 's':
			s += t.Second
		default:
			s += string(v)
		}
	}

	return s
}

func GetFormattedTime(f string) string {
	return GetTime().Format(f)
}

func CheckFileIsDirectory(path string) (bool, error)  {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if fi.IsDir() == false {
		return false, errors.New("target file is not folder")
	}

	return true, nil
}

func GetFileSize(file string) (int64, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	if fi.IsDir() {
		return 0, errors.New(fmt.Sprintf("target file %s is not file", file))
	}

	return fi.Size(), nil
}

func Chown(name string ,uid, gid int) (error) {
	err := os.Chown(name, uid, gid)
	if err != nil {
		Logger.Printf("chown %s failed", name, err)
	}

	return err
}