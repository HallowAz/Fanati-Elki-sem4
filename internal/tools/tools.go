package tools

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"
)

func StrToUint32(str string) (uint32, error) {
	res, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(res), nil
}

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, _ = rand.Read(bytes)

	return base64.URLEncoding.EncodeToString(bytes)
}
