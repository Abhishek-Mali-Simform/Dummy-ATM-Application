package services

import (
	"crypto/rand"
	"github.com/beego/beego/v2/core/logs"
	"io"
	"strconv"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func CheckError(fileName string, line int, err error) {
	if err != nil {
		logs.Error(fileName, " : line-", line, " : ", err)
	}
}

func GetRandomNumber(max int) int {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	num, _ := strconv.Atoi(string(b))
	return num
}
